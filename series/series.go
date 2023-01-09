package series

func All(n int, s string) []string {
	target := make([]string, 0)
	for n <= len(s) {
		target = append(target, s[:n])
		s = s[1:]
	}

	return target
}

func UnsafeFirst(n int, s string) string {
	if n <= len(s) {
		return s[:n]
	}

	return ""
}

func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		return s[:n], true
	}

	return "", false
}
