package group_anagrams

import (
	"strconv"
	"strings"
)

func groupAnagramsRedo(strs []string) [][]string {
	lookup := make(map[string][]string)

	for _, str := range strs {
		counter := [26]int{}

		for _, char := range str {
			counter[char-97] += 1
		}

		tokenBuilder := make([]string, len(counter))
		for index, count := range counter {
			tokenBuilder[index] = strconv.Itoa(count)
		}
		token := strings.Join(tokenBuilder, ",")

		group, ok := lookup[token]
		if ok {
			lookup[token] = append(group, str)
		} else {
			lookup[token] = []string{str}
		}
	}

	result := make([][]string, len(lookup))
	index := 0
	for _, group := range lookup {
		result[index] = group
		index += 1
	}

	return result
}
