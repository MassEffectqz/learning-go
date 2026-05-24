package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	seen := make(map[rune]bool)
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}
