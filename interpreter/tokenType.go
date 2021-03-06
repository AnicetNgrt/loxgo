package interpreter

type TokenType int

const (
	TOK_LEFT_PAREN TokenType = iota
	TOK_RIGHT_PAREN
	TOK_LEFT_BRACE
	TOK_RIGHT_BRACE
	TOK_COMMA
	TOK_DOT
	TOK_MINUS
	TOK_PLUS
	TOK_SEMICOLON
	TOK_SLASH
	TOK_STAR
	TOK_BANG
	TOK_BANG_EQUAL
	TOK_EQUAL
	TOK_EQUAL_EQUAL
	TOK_GREATER
	TOK_GREATER_EQUAL
	TOK_LESS
	TOK_LESS_EQUAL
	TOK_IDENTIFIER
	TOK_STRING
	TOK_NUMBER
	TOK_AND
	TOK_CLASS
	TOK_ELSE
	TOK_FALSE
	TOK_FUN
	TOK_FOR
	TOK_IF
	TOK_NIL
	TOK_OR
	TOK_PRINT
	TOK_RETURN
	TOK_SUPER
	TOK_THIS
	TOK_TRUE
	TOK_VAR
	TOK_WHILE
	TOK_EOF
)
