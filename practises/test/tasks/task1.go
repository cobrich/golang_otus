package tasks

import (
	"strings"
)

func TitleCase(sentence string, exclude string) string {
	sentence = strings.ToUpper(sentence)
	exclude = strings.ToUpper(exclude)

	words := strings.Split(sentence, " ")
	exclude_words := strings.Split(exclude, " ")

	for i := 0; i<len(words); i++ {
		word := words[i]
		for j := 0; j < len(exclude_words); j++ {
			ex := exclude_words[j]
			if (word  == ex) && (i > 0){
				words[i] = strings.ToLower(words[i])
			}
		}
	}

	return strings.Join(words, " ")
}