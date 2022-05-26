package maximum_depth_of_binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hamologist.com/leetcode-go/shared"
	"testing"
)

var testCases = []struct {
	root     *shared.TreeNode
	expected int
}{
	{shared.GenerateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), 3},
	{shared.GenerateTreeNode([]interface{}{1, nil, 2}), 2},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf(""),
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, maxDepth(testCase.root))
			},
		)
	}
}
