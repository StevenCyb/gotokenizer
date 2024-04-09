package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer(t *testing.T) {
	t.Parallel()

	var (
		NoneType  Type = "NONE"
		EqualType Type = "EQUAL"
		WordType  Type = "WORD"
		key            = "hello"
		separator      = "="
		value          = "world"
	)

	t.Run("Parsing", func(t *testing.T) {
		t.Parallel()

		tokenizer := New(
			fmt.Sprintf("%s%s%s", key, separator, value),
			NoneType,
			[]*Spec{
				NewSpec("^=", EqualType),
				NewSpec("^[a-z]+", WordType),
			})

		token, err := tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, WordType, token.Type)
		assert.Equal(t, key, token.Value)
		assert.Equal(t, 5, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, EqualType, token.Type)
		assert.Equal(t, separator, token.Value)
		assert.Equal(t, 6, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, WordType, token.Type)
		assert.Equal(t, value, token.Value)
		assert.Equal(t, 11, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Nil(t, token)
		assert.Equal(t, 11, tokenizer.GetCursorPosition())
	})

	t.Run("ParsingWithSkip", func(t *testing.T) {
		t.Parallel()

		var SkipType Type = "SKIP"
		tokenizer := New(
			fmt.Sprintf("  %s  %s %s ", key, separator, value),
			SkipType,
			[]*Spec{
				NewSpec(`^\s+`, SkipType),
				NewSpec("^=", EqualType),
				NewSpec("^[a-z]+", WordType),
			})

		token, err := tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, WordType, token.Type)
		assert.Equal(t, key, token.Value)
		assert.Equal(t, 7, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, EqualType, token.Type)
		assert.Equal(t, separator, token.Value)
		assert.Equal(t, 10, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, WordType, token.Type)
		assert.Equal(t, value, token.Value)
		assert.Equal(t, 16, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Nil(t, token)
		assert.Equal(t, 17, tokenizer.GetCursorPosition())
	})

	t.Run("ParsingWithUnexpected", func(t *testing.T) {
		t.Parallel()

		tokenizer := New(
			fmt.Sprintf("%s%s%s", key, separator, "#"),
			NoneType,
			[]*Spec{
				NewSpec("^=", EqualType),
				NewSpec("^[a-z]+", WordType),
			})

		token, err := tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, WordType, token.Type)
		assert.Equal(t, key, token.Value)
		assert.Equal(t, 5, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.NoError(t, err)
		assert.Equal(t, EqualType, token.Type)
		assert.Equal(t, separator, token.Value)
		assert.Equal(t, 6, tokenizer.GetCursorPosition())

		token, err = tokenizer.GetNextToken()
		assert.Error(t, err)
		assert.Nil(t, token)
	})
}
