package contains_duplicate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums     []int
	expected bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v, expected=%v", test.nums, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, containsDuplicate(test.nums))
			},
		)
	}
}
