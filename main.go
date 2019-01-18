package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	loadCSVData()
}

// EmojiCountry is a type for working with emoji country flags
// country names and country codes
type EmojiCountry struct {
	Country     string `json:"country"`
	EmojiCode   string `json:"emoji_code"`
	CountryCode string `json:"country_code"`
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
		if strings.Contains(GetEncodedUTF8StringLower(s), v.EmojiCode) {
			findings[v.Country]++
		}
	}
	return findings
}

func main() {
	log.Printf("%v", GetAmountCountryNamesMatched("Bonjour la france 🇫🇷 c'est bon aussi 🇨🇳"))
}
