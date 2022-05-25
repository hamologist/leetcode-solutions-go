package permutations

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input    []int
	expected [][]int
}{
	{
		[]int{1, 2, 3},
		[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	},
	{
		[]int{0, 1},
		[][]int{
			{0, 1},
			{1, 0},
		},
	},
	{
		[]int{1},
		[][]int{
			{1},
		},
	},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("Input: %v", testCase.input),
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, permute(testCase.input))
			},
		)
	}
}
