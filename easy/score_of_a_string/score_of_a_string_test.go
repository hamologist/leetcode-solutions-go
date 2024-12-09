package score_of_a_string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s        string
	expected int
}{
	{"hello", 13},
	{"zaz", 50},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("s=%v, expected=%v", test.s, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, scoreOfString(test.s))
			},
		)
	}
}
