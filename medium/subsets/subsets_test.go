package subsets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected [][]int
}{
	{
		[]int{1, 2, 3},
		[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
	},
	{
		[]int{0},
		[][]int{{}, {0}},
	},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("when nums is %v", testCase.nums),
			func(t *testing.T) {
				actual := subsets(testCase.nums)
				assert.Equal(t, len(testCase.expected), len(actual))
				for _, expected := range testCase.expected {
					assert.Contains(t, actual, expected)
				}
			},
		)
	}
}
