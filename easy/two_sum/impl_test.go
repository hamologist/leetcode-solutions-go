package two_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v target=%v expected=%v", test.nums, test.target, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, twoSum(test.nums, test.target))
			},
		)
	}
}
