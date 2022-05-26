package delete_node_in_a_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hamologist.com/leetcode-go/shared"
	"testing"
)

type testCase struct {
	generatingSlice []int
	head, node      *shared.ListNode
	expected        []int
}

var testCases []testCase

func Test(t *testing.T) {
	generatingSlice := []int{4, 5, 1, 9}
	testCaseOne := shared.GenerateListNode([]int{4, 5, 1, 9})
	testCaseTwo := shared.GenerateListNode([]int{4, 5, 1, 9})

	testCases = append(testCases, testCase{
		generatingSlice,
		testCaseOne,
		testCaseOne.Next,
		[]int{4, 1, 9},
	})
	testCases = append(testCases, testCase{
		generatingSlice,
		testCaseTwo,
		testCaseTwo.Next.Next,
		[]int{4, 5, 9},
	})

	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("Delete %d from %v", testCase.node.Val, generatingSlice),
			func(t *testing.T) {
				deleteNode(testCase.node)
				assert.Equal(t, testCase.expected, shared.ListToSlice(testCase.head))
			},
		)
	}
}
