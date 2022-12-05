package token

import (
	"testing"
)

func TestTokenInit(t *testing.T) {
	t := NewToken(EOF, "eof", nil, 1)
	if t == nil {
		t.Error("Failed to init a token")
	}
}
