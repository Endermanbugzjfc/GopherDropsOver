package utils

import "strings"

type word string

func (w word) Sentence() word {
	return word(strings.ToUpper(string(w[0:1]))) + w[1:]
}
