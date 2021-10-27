package main

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
	ownerURL string) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update information of owners",
		Long: `Update information of owners from remote.
Following information is available:

  Owner code
  Company
  City
  Country`,
		Example: `# Add new owners and preserve all existing owners
icm update
# Delete all owners and add most current owners
echo '{}' > $HOME/.icm/data/owner.json && icm update`,
		Args: cobra.NoArgs,
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

func parseOwners(body io.Reader) (map[string]cont.Owner, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	owners := map[string]cont.Owner{}

	var getOwnerNode func(*html.Node) error
	getOwnerNode = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.Data == "td" {

			for _, a := range n.Attr {
				if a.Key == "data-label" && a.Val == "Code" {

					codeWithU := firstChildData(n)
					if len(codeWithU) < 3 {
						return fmt.Errorf("parsing HTML failed of owner code failed because '%s' is too short", codeWithU)
					}
					code := codeWithU[0:3]
					companyNode, err := afterNextSibling(n)
					if err != nil {
						return err
					}
					cityNode, err := afterNextSibling(companyNode)
					if err != nil {
						return err
					}
					countryNode, err := afterNextSibling(cityNode)
					if err != nil {
						return err
					}
					owners[code] = cont.Owner{Code: code,
						Company: firstChildData(companyNode),
						City:    firstChildData(cityNode),
						Country: firstChildData(countryNode)}
				}
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

func afterNextSibling(n *html.Node) (*html.Node, error) {
	var next *html.Node
	if next = n.NextSibling; next != nil {
		var afterNext *html.Node
		if afterNext = next.NextSibling; afterNext != nil {
			return afterNext, nil
		}
	}
	return nil, fmt.Errorf("parsing HTML failed because nothing after next sibling of '%s'", n.Data)
}

func firstChildData(n *html.Node) string {
	var fc *html.Node
	if fc = n.FirstChild; fc != nil {
		return fc.Data
	}
	return ""
}
