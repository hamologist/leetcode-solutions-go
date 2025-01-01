package concatention_of_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
	{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v expected=%v", test.nums, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, getConcatentation(test.nums))
			},
		)
	}
}
