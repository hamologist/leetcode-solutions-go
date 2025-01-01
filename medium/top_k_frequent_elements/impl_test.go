package top_k_frequent_elements

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	k        int
	expected []int
}{
	{
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	},
	{
		[]int{1},
		1,
		[]int{1},
	},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v k=%d expected=%v", test.nums, test.k, test.expected),
			func(t *testing.T) {
				actual := topKFrequent(test.nums, test.k)
				slices.Sort(test.expected)
				slices.Sort(actual)

				assert.Equal(t, test.expected, actual)
			},
		)
	}
}
