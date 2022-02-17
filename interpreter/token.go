package interpreter

import "fmt"

type Token struct {
	tType   TokenType
	lexeme  string
	literal interface{}
	line    int
}

func NewToken(tType TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{tType, lexeme, literal, line}
}

func (t *Token) ToString() string {
	return fmt.Sprint(t.tType, " ", t.lexeme, " ", t.literal)
}
