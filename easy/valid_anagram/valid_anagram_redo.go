package valid_anagram

func isAnagramRedo(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lookup := make(map[rune]int)

	for _, char := range s {
		_, ok := lookup[char]
		if ok {
			lookup[char] += 1
		} else {
			lookup[char] = 1
		}
	}

	for _, char := range t {
		count, ok := lookup[char]
		if ok && count == 1 {
			delete(lookup, char)
		} else if ok {
			lookup[char] -= 1
		} else {
			return false
		}
	}

	return true
}
