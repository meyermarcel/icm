package owner

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"time"
	"sort"
	"os"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

const appDirectory = ".iso6346"
const updatedFileName = "updated.json"
const ownersFileName = "owners.json"
const dirPerm = 0700
const filePerm = 0644

type Owners struct {
	Owners map[string]Owner `json:"owners"`
}

type Owner struct {
	Company string `json:"company"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Updated struct {
	Updated time.Time `json:"updated"`
}

type Result struct {
	owner   Owner
	message string
}

func (mru Updated) after(updated Updated) bool {

	return mru.Updated.After(updated.Updated)
}

func NewOwner(company, city, country string) Owner {
	return Owner{company, city, country}
}

func Update() {

	homeDir, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	pathToAppDir := createPathToAppDir(homeDir)

	pathToUpdatedFile := filepath.Join(pathToAppDir, updatedFileName)
	pathToOwnersFile := filepath.Join(pathToAppDir, ownersFileName)

	remoteUpdated := getMostRecentUpdate()
	var localUpdated Updated

	file, err := ioutil.ReadFile(pathToUpdatedFile)
	if err != nil {
		writeOwners(pathToOwnersFile)
		writeUpdated(pathToUpdatedFile, Updated{time.Now()})
		fmt.Println("Updated owner defintions.")
	} else {
		json.Unmarshal([]byte(file), &localUpdated)

		if remoteUpdated.after(localUpdated) {
			writeOwners(pathToOwnersFile)
			writeUpdated(pathToUpdatedFile, Updated{time.Now()})
			fmt.Println("Owners updated.")
		} else {
			fmt.Println("Owners already up to date.")
		}
	}
}
func createPathToAppDir(homeDir string) string {

	pathToAppDir := filepath.Join(homeDir, appDirectory)

	if _, err := os.Stat(pathToAppDir); os.IsNotExist(err) {
		os.Mkdir(pathToAppDir, os.ModeDir|dirPerm)
	}

	return pathToAppDir
}

func writeUpdated(path string, updated Updated) {
	bytes, err := json.Marshal(updated)
	if err != nil {
		log.Fatalf("Could not marshal %v", updated)
	}
	ioutil.WriteFile(path, bytes, filePerm)

}

func writeOwners(path string) {
	owners := getOwners("https://www.bic-code.org/bic-letter-search/?searchterm=A")
	bytes, err := json.Marshal(owners)

	if err != nil {
		log.Fatalf("Could not marshal %v", owners)
	}
	ioutil.WriteFile(path, bytes, filePerm)
}

func getMostRecentUpdate() Updated {
	recentlyCreated := getMostRecentDate("https://www.bic-code.org/bic-codes/recently-created")
	recentlyCancelled := getMostRecentDate("https://www.bic-code.org/bic-codes/recently-cancelled")

	if recentlyCreated.After(recentlyCancelled) {
		return Updated{recentlyCreated}
	}
	return Updated{recentlyCancelled}
}

func getMostRecentDate(url string) time.Time {

	mostRecentDate, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	const query = "table tbody tr"
	dates := mostRecentDate.Find(query).FilterFunction(func(i int, selection *goquery.Selection) bool {
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
			log.Fatalf("Could not parse date '%s' because format '%s' is no longer valid", date, format)
		}
		parsedDates = append(parsedDates, parsedDate)
	}

	sort.Slice(parsedDates, func(i, j int) bool { return parsedDates[i].After(parsedDates[j]) })

	return parsedDates[0]

}

func getOwners(url string) Owners {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var owners = map[string]Owner{}

	const query = "tr td[data-label=Code]"
	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		code := s.Parent().Find("td[data-label=Code]").Text()
		company := s.Parent().Find("td[data-label=Company]").Text()
		city := s.Parent().Find("td[data-label=City]").Text()
		country := s.Parent().Find("td[data-label=Country]").Text()

		owners[code[0:3]] = NewOwner(company, city, country)
	})

	if len(owners) == 0 {
		log.Fatalf("Could not find owners in document from url '%s' with query '%s'", url, query)
	}

	return Owners{owners}
}
