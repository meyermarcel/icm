package owner

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
	"sort"
)

func Update() {

	db := InitDB()
	defer db.Close()

	if getUpdatedTimeRemote().After(getUpdatedTime(db)) {
		saveUpdatedTimeNow(db)
		deleteOwners(db)
		saveOwners(db, getOwnersRemote("https://www.bic-code.org/bic-letter-search/?searchterm=A"))
	}
}

func getUpdatedTimeRemote() time.Time {
	recentlyCreated := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-created")
	recentlyCancelled := getRecentlyDate("https://www.bic-code.org/bic-codes/recently-cancelled")

	if recentlyCreated.After(recentlyCancelled) {
		return recentlyCreated
	}
	return recentlyCancelled
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

func getOwnersRemote(url string) []Owner {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var owners []Owner

	const query = "tr td[data-label=Code]"
	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		code := s.Parent().Find("td[data-label=Code]").Text()
		company := s.Parent().Find("td[data-label=Company]").Text()
		city := s.Parent().Find("td[data-label=City]").Text()
		country := s.Parent().Find("td[data-label=Country]").Text()

		owners = append(owners, NewOwner(code[0:3], company, city, country))
	})

	if len(owners) == 0 {
		log.Fatalf("Could not find owners in document from url '%s' with query '%s'", url, query)
	}

	return owners
}
