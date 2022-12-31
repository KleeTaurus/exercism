package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) doFilter(filter func(int) bool, expected bool) Ints {
	if i == nil {
		return nil
	}

	c := Ints{}
	for _, e := range i {
		if filter(e) == expected {
			c = append(c, e)
		}
	}

	return c
}

func (i Ints) Keep(filter func(int) bool) Ints {
	return i.doFilter(filter, true)
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.doFilter(filter, false)
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	c := Lists{}
	for _, e := range l {
		if filter(e) {
			c = append(c, e)
		}
	}

	return c
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}

	c := Strings{}
	for _, e := range s {
		if filter(e) {
			c = append(c, e)
		}
	}

	return c
}
