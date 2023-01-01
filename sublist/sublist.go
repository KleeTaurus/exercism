package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	var r Relation = RelationUnequal
	switch {
	case len(l1) == len(l2):
		if isEqual(l1, l2) {
			r = RelationEqual
		}
	case len(l1) > len(l2):
		if isSuperlist(l1, l2) {
			r = RelationSuperlist
		}
	case len(l1) < len(l2):
		if isSuperlist(l2, l1) {
			r = RelationSublist
		}
	}

	return r
}

func isEqual(l1, l2 []int) bool {
	result := true
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			result = false
			break
		}
	}

	return result
}

func isSuperlist(l1, l2 []int) bool {
	if len(l2) == 0 {
		return true
	}

	result := false
	for i, e := range l1 {
		if e == l2[0] {
			subl := l1[i:]
			if len(subl) < len(l2) {
				break
			}

			if isEqual(l2, subl) {
				result = true
				break
			}
		}
	}

	return result
}
