package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrUnexpectedToken(t *testing.T) {
	t.Parallel()

	pos := 42
	key := "d"
	assert.Equal(t,
		fmt.Sprintf(errUnexpectedTokenMessage, key, pos),
		NewErrUnexpectedToken(pos, key).Error(),
	)
}
