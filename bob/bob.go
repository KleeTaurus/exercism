// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var lowers, uppers, digits int
	var question bool

	remark = strings.TrimSpace(remark)
	for i := 0; i < len(remark); i++ {
		switch {
		case remark[i] >= 'a' && remark[i] <= 'z':
			lowers++
		case remark[i] >= 'A' && remark[i] <= 'Z':
			uppers++
		case i == len(remark)-1 && remark[i] == '?':
			question = true
		case remark[i] >= '0' && remark[i] <= '9':
			digits++
		}
	}

	switch {
	case question:
		if uppers > 0 && lowers == 0 {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case uppers > 0 && lowers == 0:
		return "Whoa, chill out!"
	case uppers == 0 && lowers == 0 && digits == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
