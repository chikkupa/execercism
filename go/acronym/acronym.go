package acronym

import(
	"strings"
)

// Abbreviate returns the abbreviation of a given string
func Abbreviate(s string) string {
	var abbr []byte
	s = strings.Replace(s, "-", " ", -1)

	for _, word := range strings.Split(s, " "){
		if len(word) > 0 {
			abbr = append(abbr, strings.ToUpper(word)[0])
		}
	}

	return string(abbr)
}
