package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lookup := [26]int{}
	for i := 0; i < len(s); i++ {
		lookup[s[i]-'a'] += 1
		lookup[t[i]-'a'] -= 1
	}

	for _, check := range lookup {
		if check != 0 {
			return false
		}
	}

	return true
}
