package is_subsequence

func isSubsequence(s string, t string) bool {
	sp := 0

	for i := 0; i < len(t) && sp < len(s); i++ {
		if s[sp] == t[i] {
			sp += 1
		}
	}

	return sp == len(s)
}
