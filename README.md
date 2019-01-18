# GemojiCountries

Go package to extract country names and codes from emoji in text

For example, let's say you have the string `"Bonjour la france 🇫🇷  🇫🇷 c'est bon aussi 🇨🇳 la chine"`

The func `GetAmountCountryNamesMatched` will return `map[china:1 france:2]`
And the func `GetAmountCountryCodesMatched` will return `map[cn:1 fr:2]`

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
