package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func isAnagram(subject, candidate string) bool {
	if len(subject) != len(candidate) {
		return false
	}

	if strings.ToLower(subject) == strings.ToLower(candidate) {
		return false
	}

	subMap := strToMap(subject)
	canMap := strToMap(candidate)
	for k, v := range subMap {
		if cv, ok := canMap[unicode.ToLower(k)]; !ok || cv != v {
			return false
		}
	}

	return true
}

func strToMap(str string) map[rune]int {
	m := make(map[rune]int)
	for _, s := range str {
		m[unicode.ToLower(s)]++
	}

	return m
}
