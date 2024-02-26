package is_subsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s, t     string
	expected bool
}{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("s=%v t=%v expected=%v", test.s, test.t, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, isSubsequence(test.s, test.t))
			},
		)
	}
}
