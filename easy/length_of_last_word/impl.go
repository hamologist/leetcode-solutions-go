package length_of_last_word

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`\s+`)

func lengthOfLastWord(s string) int {
	splits := regex.Split(strings.TrimSpace(s), -1)

	return len(splits[len(splits)-1])
}
