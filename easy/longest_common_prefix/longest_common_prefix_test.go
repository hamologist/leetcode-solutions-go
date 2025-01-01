package longest_common_prefix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	strs     []string
	expected string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("%v=%v", test.strs, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, longestCommonPrefix(test.strs))
			},
		)
	}
}
