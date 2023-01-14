// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	r := regexp.MustCompile("(\\s[-\\s]*|-)")
	words := r.Split(s, -1)

	abbr := make([]byte, 0)
	for _, word := range words {
		abbr = append(abbr, getFirstCap(word))
	}
	return string(abbr)
}

func getFirstCap(word string) byte {
	for _, b := range []byte(word) {
		if b >= 65 && b <= 90 {
			return b
		}
		if b >= 97 && b <= 122 {
			return b - 32
		}
	}

	return 32
}
