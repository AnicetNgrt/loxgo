package interpreter

var KeywordToTokenType = map[string]TokenType{
	"and":    TOK_AND,
	"class":  TOK_CLASS,
	"else":   TOK_ELSE,
	"false":  TOK_FALSE,
	"for":    TOK_FOR,
	"fun":    TOK_FUN,
	"if":     TOK_IF,
	"nil":    TOK_NIL,
	"or":     TOK_OR,
	"print":  TOK_PRINT,
	"return": TOK_RETURN,
	"super":  TOK_SUPER,
	"this":   TOK_THIS,
	"true":   TOK_TRUE,
	"var":    TOK_VAR,
	"while":  TOK_WHILE,
}
