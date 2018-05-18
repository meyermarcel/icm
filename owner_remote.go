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
	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func update(pathToDB string) {

	c := make(chan time.Time)

	go getUpdatedTimeRemote(c)

	db := openDB(pathToDB)
	defer db.Close()
	updatedTimeDB := getUpdatedTime(db)

	updatedTimeRemote := <-c

	if updatedTimeRemote.After(updatedTimeDB) {
		owners := getOwnersRemote()
		deleteOwners(db)
		saveOwners(db, owners)
		saveUpdatedTimeNow(db)
	}
}

func getUpdatedTimeRemote(c chan time.Time) {
	recentlyCreated := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-created")
	recentlyCancelled := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-cancelled")

	if recentlyCreated.After(recentlyCancelled) {
		c <- recentlyCreated
	}
	c <- recentlyCancelled
}

func getRecentlyDate(url string) time.Time {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	const query = "table tbody tr"
	dates := doc.Find(query).FilterFunction(func(i int, selection *goquery.Selection) bool {
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

func getOwnersRemote() (owners []owner) {

	url := "https://www.bic-code.org/bic-letter-search/?resultsperpage=17576&searchterm="

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	const query = "tr td[data-label=Code]"
	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		code := s.Parent().Find("td[data-label=Code]").Text()
		company := s.Parent().Find("td[data-label=Company]").Text()
		city := s.Parent().Find("td[data-label=City]").Text()
		country := s.Parent().Find("td[data-label=Country]").Text()

		owners = append(owners, newOwner(code[0:3], company, city, country))
	})

	if len(owners) == 0 {
		log.Fatalf("Could not find owners in document from url '%s' with query '%s'", url, query)
	}
	return
}
