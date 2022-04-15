package main

import (
	"github.com/julienetie/basic-colly-scraper/prompt"
	"github.com/julienetie/basic-colly-scraper/scraper"
)

func main() {
	pathToScrape, selectorToQuery, fileName := prompt.Prompt(nil)
	scraper.Scraper(pathToScrape, selectorToQuery, fileName)
}
