package tokenizer

// tokenizer that lazily pulls a token from a stream.
type Tokenizer struct {
	query         string
	skipTokenType Type
	spec          []*Spec
	cursor        int
}

// GetCursorPosition return the position of the cursor.
func (t *Tokenizer) GetCursorPosition() int {
	return t.cursor
}

// HasMoreTokens checks aether we still have more tokens.
func (t *Tokenizer) HasMoreTokens() bool {
	return t.cursor < len(t.query)
}

// GetNextToken obtains next token.
func (t *Tokenizer) GetNextToken() (*Token, error) {
	if !t.HasMoreTokens() {
		return nil, nil
	}

	part := t.query[t.cursor:]

	for _, spec := range t.spec {
		matched := spec.expression.FindString(part)
		if matched == "" {
			continue
		}

		t.cursor += len(matched)
		if spec.tokenType == t.skipTokenType {
			return t.GetNextToken()
		}

		return NewToken(
			spec.tokenType,
			matched,
		), nil
	}

	return nil, NewErrUnexpectedToken(
		t.cursor,
		part[:1])
}

// New create a new tokenizer instance
// with given parameters.
func New(query string, skipTokenType Type, spec []*Spec) *Tokenizer {
	return &Tokenizer{
		cursor:        0,
		query:         query,
		skipTokenType: skipTokenType,
		spec:          spec,
	}
}
