package piglatin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var translateTests = []struct {
	in  string
	out string
}{
	{
		in:  "dog",
		out: "ogday",
	},
	{
		in:  "cat",
		out: "atcay",
	},
	{
		in:  "hat",
		out: "athay",
	},
	{
		in:  "egg",
		out: "eggday",
	},
	{
		in:  "egg hat cat dog",
		out: "eggday athay atcay ogday",
	},
}

func TestTranslate(t *testing.T) {

	for i, test := range translateTests {
		actual := Translate(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
        }

}
