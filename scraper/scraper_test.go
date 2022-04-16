package scraper

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"strconv"
	"testing"
)

const loreIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vel hendrerit odio. "

// factretriever.com/rhino-facts
// .factsList li
// test.json
// func TestScraper(t *testing.T) {
// 	// Scraper()
// 	scraper()
// }

func check(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("Err error")
	}
}

type EntryTest struct {
	Selector    string `json:selector`
	Description string `json:"description"`
}

func TestCreateJSONFile(t *testing.T) {
	fileName := "test-create-json-file.json"
	entries := make([]Entry, 0)

	for i := 1; i <= 10; i++ {
		bigI := big.NewInt(int64(i))
		entry := Entry{
			Selector:    base64.RawURLEncoding.EncodeToString(bigI.Bytes()),
			Description: loreIpsum + strconv.Itoa(i),
		}
		entries = append(entries, entry)
	}

	CreateJSONFile(&entries, &fileName)

	file, err := ioutil.ReadFile(fileName)
	check(t, err)

	data := []EntryTest{}

	err = json.Unmarshal([]byte(file), &data)
	check(t, err)

	if data[9].Selector != "Cg" {
		t.Errorf("Last entry's selector is not Cg")
	}

	if data[0].Description != loreIpsum+"1" {
		t.Errorf("First entry's description's number is not 1")
	}
}
