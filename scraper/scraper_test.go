package scraper

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"testing"
)

const loreIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis vel hendrerit odio. "

var fileName = "test-create-json-file.json"

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
		s := e.Error()
		t.Errorf(s)
	}
}

type EntryTest struct {
	Selector    string `json:selector`
	Description string `json:"description"`
}

// **CAUTION** Test will create and delete `fileName`
func TestCreateJSONFile(t *testing.T) {
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

	// Delete test file
	err = os.Remove(fileName)
	check(t, err)
}

func TestAbstractDomain(t *testing.T) {
	path := "google.com/hello-world/123"
	domain := AbstractDomain(&path)
	if domain != "google.com" {
		t.Errorf("Domain was not abstracted correctly")
	}
}
