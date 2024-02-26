package valid_anagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s, t     string
	expected bool
}{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("s=%v t=%v expected=%v", test.s, test.t, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, isAnagram(test.s, test.t))
			},
		)
	}
}
