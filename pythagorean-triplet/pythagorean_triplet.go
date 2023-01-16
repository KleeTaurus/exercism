package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	triplets := make([]Triplet, 0)
	var a, b, c int
	for a = min; a <= max-2; a++ {
		for b = a + 1; b <= max-1; b++ {
			for c = b + 1; c <= max; c++ {
				if c*c == a*a+b*b {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	triplets := make([]Triplet, 0)
	var a, b, c int
	for a = 1; ; a++ {
		b = a + 1
		c = p - a - b
		if b >= c {
			break
		}

		for b < c {
			if c*c == a*a+b*b {
				triplets = append(triplets, Triplet{a, b, c})
			}
			b++
			c--
		}
	}

	return triplets
}
