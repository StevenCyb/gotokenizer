package tokenizer

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Spec(t *testing.T) {
	t.Parallel()

	expr := "^[a-z]+"
	tokenType := Type("WORD")

	assert.Equal(t, Spec{
		expression: regexp.MustCompile(expr),
		tokenType:  tokenType,
	}, *NewSpec(expr, tokenType))
}
