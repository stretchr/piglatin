package piglatin

import (
	"strings"
)

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

func Translate(in string) string {
	var pigLatinWords []string
	englishWords := strings.Split(in, " ")

	for _, word := range englishWords {
		first := word[0:1]
		if strings.Contains(firstLetterExceptions, first) {
			pigLatinWords = append(pigLatinWords, word + firstLetterExceptionSuffix)
		} else {
			pigLatinWords = append(pigLatinWords, word[1:] + first + pigLatinSuffix)
		}
	}
	return strings.Join(pigLatinWords, " ")
}
