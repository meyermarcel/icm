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

package main

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func update() {

	ownersLastUpdateRemote := getOwnersLastUpdateRemote()

	ownersLastUpdateLocal := getOwnersLastUpdate()

	if ownersLastUpdateRemote.After(ownersLastUpdateLocal) {
		ownersRemote := getOwnersRemote()
		updateOwners(ownersRemote)
		saveNowForOwnersLastUpdate()
	}
}

func getOwnersLastUpdateRemote() time.Time {
	recentlyCreated := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-created")
	recentlyCancelled := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-cancelled")

	if recentlyCreated.After(recentlyCancelled) {
		return recentlyCreated
	}
	return recentlyCancelled
}

func getRecentlyDate(url string) time.Time {

	const query = "table tbody tr"

	dates := getBody(url).Find(query).FilterFunction(func(i int, selection *goquery.Selection) bool {
		return selection.Children().Is("td")
	}).Map(func(i int, selection *goquery.Selection) string {
		return selection.Children().First().Text()
	})

	if len(dates) == 0 {
		log.Fatalf("Could not find dates in document from url '%s' with query '%s'", url, query)
	}

	var parsedDates []time.Time

	for _, date := range dates {
		format := "2 Jan 2006"
		parsedDate, err := time.Parse(format, date)
		if err != nil {
			log.Fatalf("Could not parse date '%s' because format '%s' does not work", date, format)
		}
		parsedDates = append(parsedDates, parsedDate)
	}

	sort.Slice(parsedDates, func(i, j int) bool { return parsedDates[i].After(parsedDates[j]) })

	return parsedDates[0]
}

func getOwnersRemote() map[string]owner {
	url := "https://www.bic-code.org/bic-letter-search/?resultsperpage=17576&searchterm="

	const query = "tr td[data-label=Code]"

	owners := map[string]owner{}

	getBody(url).Find(query).Each(func(i int, s *goquery.Selection) {
		code := s.Parent().Find("td[data-label=Code]").Text()
		company := s.Parent().Find("td[data-label=Company]").Text()
		city := s.Parent().Find("td[data-label=City]").Text()
		country := s.Parent().Find("td[data-label=Country]").Text()

		owners[code[0:3]] = owner{code[0:3], company, city, country}
	})

	if len(owners) == 0 {
		log.Fatalf("Could not find owners in document from url '%s' with query '%s'", url, query)
	}
	return owners
}

func getBody(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
