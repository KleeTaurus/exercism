package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var reg = regexp.MustCompile("(\\d+)|(\\w+'?\\w+)|(\\w+)")

func WordCount(phrase string) Frequency {
	// return withRegexp(phrase)
	return withoutRegexp(phrase)
}

func withRegexp(phrase string) Frequency {
	freq := make(Frequency)
	words := reg.FindAllString(phrase, -1)
	for _, word := range words {
		freq[strings.ToLower(word)]++
	}

	return freq
}

func withoutRegexp(phrase string) Frequency {
	words := []string{}
	buf := []byte{}

	for _, c := range []byte(phrase) {
		switch {
		case c >= '0' && c <= '9':
			buf = append(buf, c)
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			buf = append(buf, c)
		case c == '\'':
			buf = append(buf, c)
		default:
			if len(buf) > 0 {
				words = append(words, string(buf))
				buf = []byte{}
			}
		}
	}
	if len(buf) > 0 {
		words = append(words, string(buf))
	}

	freq := make(Frequency)
	for _, word := range words {
		word = strings.Trim(word, "'")
		if word != "" {
			freq[strings.ToLower(word)]++
		}
	}
	return freq
}
