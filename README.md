# GemojiCountries

Go package to extract country names and codes from emoji in text

For example, let's say you have the string `"Bonjour la france 🇫🇷 c'est bon aussi 🇨🇳"`

The func `GetAmountCountryNamesMatched` will return `map[china:1 france:1]`
And the func `GetAmountCountryCodesMatched` will return `map[cn:1 fr:1]`

```go
const s = "Bonjour la france 🇫🇷 c'est bon aussi 🇨🇳"
// example with test
func TestCountries(t *testing.T) {
	t.Logf("%v", gemojicountries.GetAmountCountryNamesMatched(s))
}
```
Will print
```bash
=== RUN   TestCountriesString
--- PASS: TestCountriesString (0.00s)
	~.../utils_test.go:132: map[china:1 france:1]
PASS
ok 	0.090s
Success: Tests passed.
```