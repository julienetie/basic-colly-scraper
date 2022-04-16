package main

import (
	// "github.com/julienetie/basic-colly-scraper/prompt"
	"github.com/julienetie/basic-colly-scraper/scraper"
)

func main() {
	// pathToScrape, selectorToQuery, fileName := prompt.Prompt(nil)
	// scraper.Scraper(pathToScrape, selectorToQuery, fileName)
	//
	// #mw-content-text > div.mw-parser-output > table.wikitable.sortable.jquery-tablesorter > tbody > tr
	scraper.Scraper(
		"en.wikipedia.org/wiki/Moons_of_Jupiter",
		"tr",
		"jupiter-moons.json",
	)
}
