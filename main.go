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

func main() {
	const tSt = "Bonjour la france 🇫🇷 c'est bon aussi 🇨🇳"
	uS := strings.ToLower(fmt.Sprintf("%+q", tSt))
	findings := make(map[string]int)
	for _, v := range EmojiCountryData {
		if strings.Contains(uS, v.EmojiCode) {
			findings[v.CountryCode]++
		}
	}
	log.Printf("%v", findings)
}
