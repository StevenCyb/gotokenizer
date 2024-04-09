package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Token(t *testing.T) {
	t.Parallel()

	tokenType := Type("WORD")
	value := "hello"
	token := NewToken(tokenType, value)

	assert.Equal(t, Token{
		Type:  tokenType,
		Value: value,
	}, *token)
	assert.Equal(t, string(tokenType), token.Type.String())
}
