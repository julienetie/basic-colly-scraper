package prompt

import (
	"bufio"
	"fmt"
	"os"
)

type FakeInput struct {
	x string
	y string
	z string
}

func isInputValid(input string) (bool, string) {
	if len(input) < 2 {
		return false, "Input too short"
	}

	return true, ""
}

func setPath(pathToScrape *string) {
	fmt.Printf(`Input the path you want to scrape "e.g https://en.wikipedia.org/wiki/Moons_of_Jupiter": `)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*pathToScrape = scanner.Text()
	}
	// fmt.Scanln(pathToScrape)
	isValid, message := isInputValid(*pathToScrape)

	if isValid != true {
		fmt.Println(message)
		setPath(pathToScrape)
	}
}

func setSelector(selectorToQuery *string) {
	fmt.Printf(`Define the selector to query (XPath or CSS selector): `)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*selectorToQuery = scanner.Text()
	}
	// fmt.Scanln(pathToScrape)
	isValid, message := isInputValid(*selectorToQuery)

	if isValid != true {
		fmt.Println(message)
		setSelector(selectorToQuery)
	}
}

func setFileName(fileName *string) {
	fmt.Printf(`Save the JSON file as: `)
	fmt.Scanln(fileName)
	isValid, message := isInputValid(*fileName)

	if isValid != true {
		fmt.Println(message)
		setFileName(fileName)
	}
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

	setPath(&pathToScrape)
	setSelector(&selectorToQuery)
	setFileName(&fileName)

	return pathToScrape, selectorToQuery, fileName
}
