package interpreter

type Parser struct {
	tokens  []Token
	current int
	errors  []Error
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
		errors:  make([]Error, 0),
	}
}

func (p *Parser) Parse() (Expr, []Error) {
	expr, _ := p.expression()
	return expr, p.errors
}

func (p *Parser) expression() (Expr, bool) {
	return p.equality()
}

func (p *Parser) equality() (Expr, bool) {
	expr, ok := p.comparison()
	if !ok {
		return nil, false
	}

	for p.match(TOK_BANG_EQUAL, TOK_EQUAL_EQUAL) {
		operator := p.previous()
		right, ok := p.comparison()
		if !ok {
			return nil, false
		}
		expr = &Binary{expr, operator, right}
	}

	return expr, true
}

func (p *Parser) comparison() (Expr, bool) {
	expr, ok := p.term()
	if !ok {
		return nil, false
	}

	for p.match(TOK_GREATER, TOK_GREATER, TOK_LESS, TOK_LESS_EQUAL, TOK_GREATER_EQUAL) {
		operator := p.previous()
		right, ok := p.term()
		if !ok {
			return nil, false
		}
		expr = &Binary{expr, operator, right}
	}

	return expr, true
}

func (p *Parser) term() (Expr, bool) {
	expr, ok := p.factor()
	if !ok {
		return nil, false
	}

	for p.match(TOK_MINUS, TOK_PLUS) {
		operator := p.previous()
		right, ok := p.factor()
		if !ok {
			return nil, false
		}
		expr = &Binary{expr, operator, right}
	}

	return expr, true
}

func (p *Parser) factor() (Expr, bool) {
	expr, ok := p.unary()
	if !ok {
		return nil, false
	}

	for p.match(TOK_SLASH, TOK_STAR) {
		operator := p.previous()
		right, ok := p.unary()
		if !ok {
			return nil, false
		}
		expr = &Binary{expr, operator, right}
	}

	return expr, true
}

func (p *Parser) unary() (Expr, bool) {
	if p.match(TOK_BANG, TOK_MINUS) {
		operator := p.previous()
		right, ok := p.unary()
		return &Unary{operator, right}, ok
	}
	return p.primary()
}

func (p *Parser) primary() (Expr, bool) {
	if p.match(TOK_FALSE) {
		return &Literal{false}, true
	}
	if p.match(TOK_TRUE) {
		return &Literal{true}, true
	}
	if p.match(TOK_NIL) {
		return &Literal{nil}, true
	}

	if p.match(TOK_NUMBER, TOK_STRING) {
		return &Literal{p.previous().literal}, true
	}

	if p.match(TOK_LEFT_PAREN) {
		expr, okExpr := p.expression()
		_, okParenth := p.consume(TOK_RIGHT_PAREN, "Expect ')' after expression.")
		return &Grouping{expr}, okExpr && okParenth
	}

	p.error(p.peek(), "Expected expression.")
	return nil, false
}

func (p *Parser) consume(tType TokenType, msg string) (Token, bool) {
	if p.check(tType) {
		return p.advance(), true
	}

	token := p.peek()
	p.error(token, msg)
	return token, false
}

func (p *Parser) error(token Token, msg string) {
	err := ERR(token.line, msg)
	p.errors = append(p.errors, err)
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().tType == TOK_SEMICOLON {
			return
		}
		switch p.peek().tType {
		case TOK_CLASS, TOK_FOR, TOK_FUN, TOK_IF, TOK_PRINT, TOK_RETURN, TOK_VAR, TOK_WHILE:
			return
		}
	}
	p.advance()
}

func (p *Parser) match(types ...TokenType) bool {
	for _, tType := range types {
		if p.check(tType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().tType == tType
}

func (p *Parser) isAtEnd() bool {
	return p.peek().tType == TOK_EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}
