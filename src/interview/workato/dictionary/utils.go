package dictionary

import "strings"

//var words []string = []string{"hello","wix","bye","bay","bucks"}
var words []string = []string{"hello", "wix", "bye", "bucks"}

func IsWord(s string) int8 {
	for _, w := range words {
		if s == w {
			return 1
		}
		if strings.HasPrefix(w, s) {
			return -1
		}
	}
	return 0
}
