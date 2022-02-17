package interpreter

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(TOK_BANG_EQUAL, TOK_EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for p.match(TOK_GREATER, TOK_GREATER, TOK_LESS, TOK_LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(TOK_MINUS, TOK_PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for p.match(TOK_SLASH, TOK_STAR) {
		operator := p.previous()
		right := p.unary()
		expr = &Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(TOK_BANG, TOK_MINUS) {
		operator := p.previous()
		right := p.unary()
		return &Unary{operator, right}
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(TOK_FALSE) {
		return &Literal{false}
	}
	if p.match(TOK_TRUE) {
		return &Literal{true}
	}
	if p.match(TOK_NIL) {
		return &Literal{nil}
	}

	if p.match(TOK_NUMBER, TOK_STRING) {
		return &Literal{p.previous().literal}
	}

	if p.match(TOK_LEFT_PAREN) {
		expr := p.expression()
		p.consume(TOK_RIGHT_PAREN, "Expect ')' after expression.")
		return &Grouping{expr}
	}

	return &Literal{nil} // TODO Change this
}

func (p *Parser) consume(tType TokenType, msg string) {

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
