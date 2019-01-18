package gemojicountries

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"index/suffixarray"
	"io"
	"os"
	"regexp"
	"strings"
)

func init() {
	loadCSVData()
}

// EmojiCountry is a type for working with emoji country flags
// country names and country codes
type EmojiCountry struct {
	Country     string
	EmojiRegex  *regexp.Regexp
	EmojiCode   string
	CountryCode string
}

// EmojiCountryData is a slice of EmojiCountry
var EmojiCountryData []EmojiCountry

func loadCSVData() {
	f, _ := os.Open("./emoji_countries.csv")
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		EmojiCountryData = append(EmojiCountryData, EmojiCountry{
			Country:     strings.ToLower(record[0]),
			EmojiRegex:  regexp.MustCompile(strings.Replace(strings.ToLower(record[1]), `\`, `\\`, -1)),
			EmojiCode:   strings.ToLower(record[1]),
			CountryCode: strings.ToLower(record[2]),
		})
	}
}

// GetEncodedUTF8StringLower returns a string encoded in UTF8, i.e. the code
// points in lower case for using with the emoji_countries.csv file, which is in
// upper case, as the ones generated from several emoji libraries and used
// in other programming languages, in Go it's lower case.
func GetEncodedUTF8StringLower(s string) string {
	return strings.ToLower(fmt.Sprintf("%+q", s))
}

// GetAmountCountryNamesMatched returns the names of the countries matched
// and the amount of each
func GetAmountCountryNamesMatched(s string) map[string]int {
	findings := make(map[string]int)
	for _, v := range EmojiCountryData {
		index := suffixarray.New([]byte(GetEncodedUTF8StringLower(s)))
		results := index.FindAllIndex(v.EmojiRegex, -1)
		if strings.Contains(GetEncodedUTF8StringLower(s), v.EmojiCode) {
			findings[v.Country] += len(results)
		}
	}
	return findings
}

// GetAmountCountryCodesMatched returns the countries matched represented by
// their country codes and the amount of each
func GetAmountCountryCodesMatched(s string) map[string]int {
	findings := make(map[string]int)
	for _, v := range EmojiCountryData {
		index := suffixarray.New([]byte(GetEncodedUTF8StringLower(s)))
		results := index.FindAllIndex(v.EmojiRegex, -1)
		if strings.Contains(GetEncodedUTF8StringLower(s), v.EmojiCode) {
			findings[v.CountryCode] += len(results)
		}
	}
	return findings
}
