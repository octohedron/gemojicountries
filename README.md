# GemojiCountries

Go package to extract country names and codes from emoji in text


### Install

+ Make sure you have your `$GOPATH` set, otherwise 

```bash
$ export GOPATH=~/go # Or your gopath location
$ go get -u github.com/octohedron/gemojicountries
```

### Overview
For example, let's say you have the string `"Bonjour la france 🇫🇷  🇫🇷 c'est bon, aussi la chine 🇨🇳"`

+ `GetAmountCountryNamesMatched` will return `map[china:1 france:2]`
+ `GetAmountCountryCodesMatched` will return `map[cn:1 fr:2]`

Note that it **only matches emoji**, not the country names

If you want to get the country codes and you have the country names, i.e. `"russia"`, you can also use the util func

+ `GetCountryCodeByCountry` will return `"ru"`

For getting a string with flag emoji with the flags replaced with their string format i.e. `"Bonjour la 🇫🇷 "`

+ `GetStringReplacedEmoji` will return `"Bonjour la france"`

Example with test

```go
package main

import (
	"testing"

	"github.com/octohedron/gemojicountries"
)

// example with test
func TestCountriesString(t *testing.T) {
	const s = "Bonjour la france 🇫🇷  🇫🇷 c'est bon aussi 🇨🇳 la chine"
	t.Logf("%v", gemojicountries.GetAmountCountryCodesMatched(s))
}
```

Will print

```bash
=== RUN   TestCountriesString
--- PASS: TestCountriesString (0.00s)
	../utils_test.go:132: map[china:1 france:2]
PASS
ok 	0.090s
Success: Tests passed.
```

Contributing? Adding countries and features? Improving performance? Yes. Send PRs 👍
