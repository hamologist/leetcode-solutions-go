package remove_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	nums         []int
	val          int
	expectedK    int
	expectedNums []int
}{
	{
		[]int{3, 2, 2, 3},
		3,
		2,
		[]int{2, 2},
	},
	{
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		5,
		[]int{0, 1, 3, 0, 4},
	},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("nums=%v expectedK=%d expectedNums=%v", test.nums, test.expectedK, test.expectedNums),
			func(t *testing.T) {
				actual := removeElement(test.nums, test.val)
				assert.Equal(t, test.expectedK, actual)
				assert.Equal(t, test.expectedNums, test.nums[:actual])
			},
		)
	}
}
