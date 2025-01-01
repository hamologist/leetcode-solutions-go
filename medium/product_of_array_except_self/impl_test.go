package product_of_array_except_self

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	},
	{
		[]int{-1, 1, 0, -3, 3},
		[]int{0, 0, 9, 0, 0},
	},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v expected=%v", test.nums, test.expected),
			func(t *testing.T) {
				actual := productExceptSelf(test.nums)
				assert.Equal(t, test.expected, actual)
			},
		)
	}
}
