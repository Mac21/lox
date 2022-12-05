package token

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}

func NewToken(t TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		tokenType: t,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}
