package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Entry struct {
	Selector    string `json:selector`
	Description string `json:"description"`
}

func Check(e error, message string) {
	if e != nil {
		if message != "" {
			log.Println(message)
		}
		panic(e)
	}
}

func CreateJSONFile(data *[]Entry, fileName *string) {
	file, err := json.MarshalIndent(data, "", " ")
	Check(err, "Unable to create JSON file")
	_ = ioutil.WriteFile(*fileName, file, 0644)
}

func WriteStdout(entries *[]Entry) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", " ")
	encoder.Encode(entries)
}

func abstractDomain(pathToScrape *string) string {
	domainLastIndex := strings.Index(*pathToScrape, "/")
	if domainLastIndex > 0 {
		return (*pathToScrape)[0:domainLastIndex]
	}
	return *pathToScrape
}

func Scraper(pathToScrape, selectorToQuery, fileName string) {
	var domain string
	entries := make([]Entry, 0)

	// Get the domain from the path
	domain = abstractDomain(&pathToScrape)

	// Setup scraper
	collector := colly.NewCollector(
		colly.AllowedDomains(domain, "www."+domain),
	)
	collector.OnHTML(selectorToQuery, func(element *colly.HTMLElement) {
		factDesc := element.Text

		entry := Entry{
			Selector:    selectorToQuery,
			Description: factDesc,
		}

		entries = append(entries, entry)
	})

	// Notifiy when page request starts
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit("https://" + pathToScrape)

	// Write to file
	CreateJSONFile(&entries, &fileName)

	// Write to stdout
	WriteStdout(&entries)
}
