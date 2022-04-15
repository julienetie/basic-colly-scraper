package scraper

import (
	"github.com/gocolly/colly"
	"strconv"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Scraper() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		Check(err)

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")
}
