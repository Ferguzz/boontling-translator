package boontling

import (
	"testing"
)

func TestTranslateSingleWord(t *testing.T) {
	result, _ := translate("good")
	if result != "bahl" {
		t.Error(result)
	}
}

func TestTranslateReturnsErrorIfResultIsSame(t *testing.T) {
	result, err := translate("foo bar baz")
	if err == nil {
		t.Error(result)
	}
}

func TestPhrases(t *testing.T) {
	result, err := translate("lots")
	if err == nil {
		t.Error(result)
	}

	result, _ = translate("lots drink")
	if result != "lots horn" {
		t.Error(result)
	}

	result, _ = translate("lots of")
	if result != "heelch" {
		t.Error(result)
	}

	result, _ = translate("lots of nothing")
	if result != "heelch nothing" {
		t.Error(result)
	}

	result, _ = translate("lots of boobs")
	if result != "heelch Mollies" {
		t.Error(result)
	}

	result, _ = translate("twenty five drinks")
	if result != "twenty five horns" {
		t.Error(result)
	}

	result, _ = translate("twenty five cents")
	if result != "toobs" {
		t.Error(result)
	}
}

func TestGetWordPlurals(t *testing.T) {
	result, _ := getWord("boob", nouns)
	if result != "Mollie" {
		t.Error(result)
	}

	result, _ = getWord("boobs", nouns)
	if result != "Mollies" {
		t.Error(result)
	}

	result, _ = getWord("boobies", nouns)
	if result != "Mollies" {
		t.Error(result)
	}
}
