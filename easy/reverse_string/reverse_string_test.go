package reverse_string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	input, expected []byte
}{
	{[]byte("hello"), []byte("olleh")},
	{[]byte("Hannah"), []byte("hannaH")},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("Input: \"%s\", Output: \"%s\"", test.input, test.expected),
			func(t *testing.T) {
				reverseString(test.input)
				assert.Equal(t, test.expected, test.input)
			},
		)
	}
}
