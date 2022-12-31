package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newFormat := make(map[string]int)
	for k, v := range in {
		for _, l := range v {
			newFormat[strings.ToLower(l)] = k
		}
	}

	return newFormat
}
