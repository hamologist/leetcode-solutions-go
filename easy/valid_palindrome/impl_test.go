package valid_palindrome

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input    string
	expected bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("input=%v expected=%v", test.input, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, isPalindrome(test.input))
			},
		)
	}
}
