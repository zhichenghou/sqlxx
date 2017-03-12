package engin

import (
	"strings"
	"unicode"
)

// for example: transfer browse_by_set to BrowseBySet
func UnderscoreToCamelcase(str string, title bool) string {
	return LowerToCamelcase(str, "_", title)
}

// for example: transfer BrowseBySet to browse_by_set
func CamelcaseToUnderscore(str string) string {
	return CamelcaseToLower(str, "_")
}

func LowerToCamelcase(str string, sp string, title bool) string {
	var method string
	sli := strings.Split(str, sp)
	for i, v := range sli {
		if i == 0 {
			if title {
				method += strings.Title(v)
			} else {
				method += v
			}
		} else {
			method += strings.Title(v)
		}
	}
	return method
}

func CamelcaseToLower(str string, sp string) string {
	return strings.Join(CamelcaseToSlice(str, true), sp)
}

func CamelcaseToSlice(str string, toLower bool) []string {
	var words []string
	l := 0
	i := 1

	for s := str; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l < 1 {
			l = len(s)
		}

		if toLower {
			words = append(words, strings.ToLower(s[:l]))
		} else {
			words = append(words, s[:l])
		}

		i++
	}

	return words
}
