package scanner

import (
	"github.com/mac21/lox/token"
)

type Scanner struct {
	start   int
	current int
	line    int
	source  string
	tokens  []*token.Token
}

func (s Scanner) scanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		start:   0,
		current: 0,
		line:    1,
		source:  source,
		tokens:  make([]*token.Token, 0),
	}
}
