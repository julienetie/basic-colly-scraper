package prompt

import (
	"fmt"
	// "os"
)

type FakeInput struct {
	x string
	y string
	z string
}

// #mw-content-text > div.mw-parser-output > table.wikitable.sortable.jquery-tablesorter > tbody > tr
func Prompt(fakeInput *FakeInput) (string, string, string) {
	var pathToScrape string
	var selectorToQuery string
	var fileName string

	// This will not test the stdin
	if fakeInput != nil {
		return fakeInput.x, fakeInput.y, fakeInput.z
	}

	fmt.Printf(`Input the path you want to scrape "e.g https://en.wikipedia.org/wiki/Moons_of_Jupiter": `)
	fmt.Scanln(&pathToScrape)

	fmt.Printf(`Define the selector to query (XPath or CSS selector): `)
	fmt.Scanln(&selectorToQuery)

	fmt.Printf(`Save the JSON file as: `)
	fmt.Scanln(&fileName)

	return pathToScrape, selectorToQuery, fileName
}
