package add_two_integers

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

var testCases = []struct {
	num1, num2, expected int
}{
	{12, 5, 17},
	{-10, 4, -6},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		t.Run(
			fmt.Sprintf("%d+%d=%d", test.num1, test.num2, test.expected),
			func(t *testing.T) {
				assert.Equal(t, test.expected, sum(test.num1, test.num2))
			},
		)
	}
}
