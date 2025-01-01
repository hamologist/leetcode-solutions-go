package group_anagrams

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	strs     []string
	expected [][]string
}{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
	{
		[]string{""},
		[][]string{{""}},
	},
	{
		[]string{"a"},
		[][]string{{"a"}},
	},
}

func less(s [][]string) func(i, j int) bool {
	return func(i, j int) bool {
		return len(s[i]) < len(s[j])
	}
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("strs=%v expected=%v", test.strs, test.expected),
			func(t *testing.T) {
				actual := groupAnagrams(test.strs)

				for _, strs := range test.expected {
					slices.Sort(strs)
				}
				sort.Slice(test.expected, less(test.expected))

				for _, strs := range actual {
					slices.Sort(strs)
				}
				sort.Slice(actual, less(actual))

				assert.Equal(t, test.expected, actual)
			},
		)
	}
}

func TestRedo(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("strs=%v expected=%v", test.strs, test.expected),
			func(t *testing.T) {
				actual := groupAnagramsRedo(test.strs)

				for _, strs := range test.expected {
					slices.Sort(strs)
				}
				sort.Slice(test.expected, less(test.expected))

				for _, strs := range actual {
					slices.Sort(strs)
				}
				sort.Slice(actual, less(actual))

				assert.Equal(t, test.expected, actual)
			},
		)
	}
}
