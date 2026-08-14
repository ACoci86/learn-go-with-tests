package iteration

import "strings"

const repeatedCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatedCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
