package main

import "fmt"

type Token struct {
	tokType TokenType
	lexeme  string
	literal *Object
	line    int
}

func NewToken(tokType TokenType, lexeme string, literal *Object, line int) *Token {
	return &Token{tokType, lexeme, literal, line}
}

func (t *Token) toString() string {
	return fmt.Sprint(t.tokType, " ", t.lexeme, " ", t.literal)
}
