package cmd

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"

	"github.com/meyermarcel/icm/cont"
	"github.com/meyermarcel/icm/data"
	"github.com/spf13/cobra"
)

func newUpdateOwnerCmd(
	ownerUpdater data.OwnerUpdater,
	timestampUpdater data.TimestampUpdater,
	ownerURL string,
) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update information of owners",
		Long: `Update information of owners from remote.
Following information is available:

  Owner code
  Company
  City
  Country`,
		Example: `# Overwrite owner.csv file with newest owners
icm update
# Create custom-owner.csv to have additional custom mapping of owner codes
# Use semicolon as a separator. For using double quotes please see existing
# owner.csv file.
echo 'AAA;my company;my city;my country' >> $HOME/.icm/data/custom-owner.csv`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(ownerUpdater, timestampUpdater, ownerURL)
		},
	}
	return updateCmd
}

func update(ownerUpdater data.OwnerUpdater, timestampUpdater data.TimestampUpdater, ownerURL string) error {
	if err := timestampUpdater.Update(); err != nil {
		return err
	}

	resp, err := http.Get(ownerURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	owners, err := parseOwners(resp.Body)
	if err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return ownerUpdater.Update(owners)
}

func parseOwners(body io.Reader) ([]cont.Owner, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	owners := make([]cont.Owner, 0)

	var getOwnerNode func(*html.Node) error
	getOwnerNode = func(n *html.Node) error {
		codeWithU := parseHTMLtdData(n, "Code:")
		if codeWithU != "" {

			if len(codeWithU) < 4 {
				return fmt.Errorf("parsing HTML failed of owner code failed because '%s' is too short", codeWithU)
			}
			code := codeWithU[0:3]

			companyTdNode := nextTdSibling(n)
			company := parseHTMLtdData(companyTdNode, "Company:")
			cityNode := nextTdSibling(companyTdNode)
			city := parseHTMLtdData(cityNode, "City:")
			countryNode := nextTdSibling(cityNode)
			country := parseHTMLtdData(countryNode, "Country:")

			owners = append(owners,
				cont.Owner{
					Code:    code,
					Company: company,
					City:    city,
					Country: country,
				},
			)

			// If valid td tag found continue with sibling
			// instead of parsing every child.
			err := getOwnerNode(n.NextSibling)
			if err != nil {
				return err
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := getOwnerNode(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = getOwnerNode(doc)
	if err != nil {
		return nil, err
	}
	if len(owners) == 0 {
		return nil, fmt.Errorf("parsing HTML failed because no owner was parsed")
	}
	return owners, nil
}

func parseHTMLtdData(td *html.Node, spanDescription string) string {
	if td.Type == html.ElementNode && td.Data == "td" {
		// Iterate through nodes because simple inner text (e.g. '\n') is also a node.
		for c := td.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "span" && c.FirstChild != nil && c.FirstChild.Data == spanDescription {
				for ns := c.NextSibling; ns != nil; ns = ns.NextSibling {
					if ns.Type == html.ElementNode && ns.Data == "span" {
						// Empty span:
						// <span></span>
						if ns.FirstChild == nil {
							return ""
						}
						return ns.FirstChild.Data
					}
				}
			}
		}
	}
	return ""
}

func nextTdSibling(td *html.Node) *html.Node {
	for ns := td.NextSibling; ns != nil; ns = ns.NextSibling {
		if ns.Type == html.ElementNode && ns.Data == "td" {
			return ns
		}
	}
	return nil
}
