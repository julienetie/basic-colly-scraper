package main

import (
	"github.com/julienetie/basic-colly-scraper/prompt"
	// "github.com/julienetie/basic-colly-scraper/scraper"
	"fmt"
	"reflect"
)

func main() {
	pathToScrape, selectorToQuery, fileName := prompt.Prompt(nil)
	// scraper.Scraper()
	fmt.Println(len(pathToScrape), reflect.TypeOf(selectorToQuery), fileName)
}
