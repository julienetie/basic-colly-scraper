package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fact struct {
	ID          int    `json:"id"`
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

func writeJSON(data []Fact, fileName *string) {
	file, err := json.MarshalIndent(data, "", " ")
	Check(err, "Unable to create JSON file")

	_ = ioutil.WriteFile(*fileName, file, 0644)
}

// factretriever.com/rhino-facts
// .factsList li
// test.json
func Scraper(pathToScrape, selectorToQuery, fileName string) {
	var domain string
	domainLastIndex := strings.Index(pathToScrape, "/")

	if domainLastIndex <= 0 {
		domain = pathToScrape
	} else {
		domain = pathToScrape[0:domainLastIndex]
	}

	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains(domain, "www."+domain),
	)

	fmt.Println(selectorToQuery)
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		Check(err, "")

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://" + pathToScrape)

	writeJSON(allFacts, &fileName)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", " ")
	encoder.Encode(allFacts)
}
