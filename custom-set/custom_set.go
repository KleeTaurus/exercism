package stringset

import (
	"bytes"
	"fmt"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	m map[string]bool
}

func New() Set {
	return Set{
		m: map[string]bool{},
	}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, k := range l {
		s.m[k] = true
	}

	return s
}

func (s Set) String() string {
	var b bytes.Buffer
	for k := range s.m {
		b.WriteString(fmt.Sprintf("\"%s\", ", k))
	}

	if b.Len() > 2 {
		// remove tail ", "
		b.Truncate(b.Len() - 2)
	}

	return "{" + b.String() + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.m) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.m[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.m[elem] = true
}

func Subset(s1, s2 Set) bool {
	is := true
	for k := range s1.m {
		_, ok := s2.m[k]
		if !ok {
			is = false
			break
		}
	}

	return is
}

func Disjoint(s1, s2 Set) bool {
	insec := Intersection(s1, s2)
	return len(insec.m) == 0
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && len(s1.m) == len(s2.m)
}

func Intersection(s1, s2 Set) Set {
	var small, large Set
	if len(s1.m) < len(s2.m) {
		small, large = s1, s2
	} else {
		small, large = s2, s1
	}

	target := New()
	for k := range small.m {
		if _, ok := large.m[k]; ok {
			target.Add(k)
		}
	}

	return target
}

func Difference(s1, s2 Set) Set {
	target := New()
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			target.Add(k)
		}
	}

	return target
}

func Union(s1, s2 Set) Set {
	target := New()
	for k := range s1.m {
		target.Add(k)
	}
	for k := range s2.m {
		target.Add(k)
	}

	return target
}
