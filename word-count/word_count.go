package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var reg = regexp.MustCompile("(\\d+)|(\\w+'?\\w+)|(\\w+)")

func WordCount(phrase string) Frequency {
	return withRegexp(phrase)
}

func withRegexp(phrase string) Frequency {
	freq := make(Frequency)
	words := reg.FindAllString(phrase, -1)
	for _, word := range words {
		freq[strings.ToLower(word)]++
	}

	return freq
}
