package majority_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected int
}{
	{[]int{3, 2, 3}, 3},
	{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v expected=%d", test.nums, test.expected),
			func(t *testing.T) {
				actual := majorityElement(test.nums)
				assert.Equal(t, test.expected, actual)
			},
		)
	}
}
