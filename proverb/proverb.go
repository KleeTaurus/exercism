// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	fullText := []string{}
	for i := 0; i < len(rhyme); i++ {
		if i == len(rhyme)-1 {
			fullText = append(fullText, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			fullText = append(fullText, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	return fullText
}
