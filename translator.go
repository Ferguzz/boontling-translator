package boontling

import (
	"errors"
	"strings"
)

var pronouns []string = []string{"i", "you", "he", "she", "it", "we", "they", "let's"}
var pluralSuffixes []string = []string{"s", "es", "ies"}

func isPronoun(s string) bool {
	for _, pronoun := range pronouns {
		if s == pronoun {
			return true
		}
	}
	return false
}

func getWord(word string, dictionary map[string]string) (string, bool) {
	// deal with 'ing' verbs too

	boontling, exists := dictionary[word]
	if !exists && strings.HasSuffix(word, "s") {
		for _, suffix := range pluralSuffixes {
			if strings.HasSuffix(word, suffix) {
				boontling, exists = dictionary[word[:len(word)-len(suffix)]]
				if exists {
					boontling = strings.Join([]string{boontling, "s"}, "")
					break
				}
			}
		}
	}
	return boontling, exists
}

func translate(input string) (string, error) {
	words := strings.Split(input, " ")
	buffer := make([]string, 0, len(words))

	for i := 0; i < len(words); i++ {
		word := words[i]
		boontling, exists := getWord(word, nouns)
		next_word_index := i
		for exists {
			if !strings.HasPrefix(boontling, "~") {
				// Not a continuing phrase.
				break
			}

			if next_word_index == len(words)-1 {
				// Last word.
				boontling, exists = word, false
				break
			}

			if words[next_word_index+1] != boontling[1:] {
				// Next word doesn't match.
				boontling, exists = word, false
				break
			}

			boontling, exists = getWord(boontling, nouns)
			next_word_index++
		}

		if exists {
			i = next_word_index
			buffer = append(buffer, boontling)
		} else {
			buffer = append(buffer, word)
		}
	}

	result := strings.Join(buffer, " ")
	if result == input {
		return "", errors.New("Input not translated.")
	}
	return result, nil
}
