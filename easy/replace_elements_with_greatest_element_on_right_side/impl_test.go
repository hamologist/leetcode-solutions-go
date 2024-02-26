package replace_elements_with_greatest_element_on_right_side

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
	{[]int{400}, []int{-1}},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v expected=%v", test.nums, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, replaceElements(test.nums))
			},
		)
	}
}
