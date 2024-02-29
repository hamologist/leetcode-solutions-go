package length_of_last_word

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s        string
	expected int
}{
	{"Hello World", 5},
	{"   fly me   to   the moon  ", 4},
	{"luffy is still joyboy", 6},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("s=%v expected=%v", test.s, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, lengthOfLastWord(test.s))
			},
		)
	}
}
