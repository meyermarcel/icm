// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/meyermarcel/icm/internal/cont"
	"github.com/meyermarcel/icm/internal/data"
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
	owners, err := ownersRemote(ownerURL)
	if err != nil {
		return err
	}
	if err := ownerUpdater.Update(owners); err != nil {
		return err
	}
	return nil
}

func ownersRemote(url string) (map[string]cont.Owner, error) {

	const query = "tr td[data-label=Code]"

	owners := map[string]cont.Owner{}

	document, err := getBody(url)
	if err != nil {
		return nil, err
	}
	document.Find(query).Each(func(i int, s *goquery.Selection) {
		code := s.Parent().Find("td[data-label=Code]").Text()
		company := s.Parent().Find("td[data-label=Company]").Text()
		city := s.Parent().Find("td[data-label=City]").Text()
		country := s.Parent().Find("td[data-label=Country]").Text()

		owners[code[0:3]] = cont.Owner{Code: code[0:3], Company: company, City: city, Country: country}
	})

	if len(owners) == 0 {
		return nil, fmt.Errorf("could not find owners in document from url '%s' with query '%s'", url, query)
	}
	return owners, nil
}

func getBody(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
