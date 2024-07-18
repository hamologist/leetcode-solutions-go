package valid_palindrome

import (
	"strings"
)

func strip(s string) string {
	s = strings.ToLower(s)
	var builder strings.Builder
	for _, char := range s {
		if ('a' <= char && 'z' >= char) ||
			('0' <= char && '9' >= char) {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func isPalindrome(s string) bool {
	s = strip(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
