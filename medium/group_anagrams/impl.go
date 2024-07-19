package group_anagrams

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		var lookup [26]int

		for _, char := range s {
			lookup[char-97] += 1
		}

		var builder strings.Builder
		for _, count := range lookup {
			builder.WriteString(fmt.Sprintf("%v,", strconv.Itoa(count)))
		}
		key := builder.String()

		_, ok := groups[key]
		if ok {
			groups[key] = append(groups[key], s)
		} else {
			groups[key] = []string{s}
		}
	}

	var result [][]string
	for _, value := range groups {
		result = append(result, value)
	}
	return result
}
