package main

type Scanner struct {
	source  []rune
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{[]rune(source), make([]Token, 0), 0, 0, 1}
}

func (s *Scanner) scanTokens() ([]Token, []Error) {
	errors := make([]Error, 0)

	for {
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		err := s.scanToken()

		if err != nil {
			errors = append(errors, err)
		}
	}

	s.tokens = append(s.tokens, *NewToken(TOK_EOF, "", nil, s.line))

	return s.tokens, errors
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() Error {
	token := s.advance()

	switch token {
	case "(":
		s.addToken(TOK_LEFT_PAREN)
	case ")":
		s.addToken(TOK_RIGHT_PAREN)
	case "{":
		s.addToken(TOK_LEFT_BRACE)
	case "}":
		s.addToken(TOK_RIGHT_PAREN)
	case ",":
		s.addToken(TOK_COMMA)
	case ".":
		s.addToken(TOK_DOT)
	case "-":
		s.addToken(TOK_MINUS)
	case "+":
		s.addToken(TOK_PLUS)
	case ";":
		s.addToken(TOK_SEMICOLON)
	case "*":
		s.addToken(TOK_STAR)
	case "!":
		if s.match("=") {
			s.addToken(TOK_BANG_EQUAL)
		} else {
			s.addToken(TOK_BANG)
		}
	case "=":
		if s.match("=") {
			s.addToken(TOK_EQUAL_EQUAL)
		} else {
			s.addToken(TOK_EQUAL)
		}
	case "<":
		if s.match("=") {
			s.addToken(TOK_LESS_EQUAL)
		} else {
			s.addToken(TOK_LESS)
		}
	case ">":
		if s.match("=") {
			s.addToken(TOK_GREATER_EQUAL)
		} else {
			s.addToken(TOK_GREATER)
		}
	case "/":
		if s.match("/") {

		} else {
			s.addToken(TOK_SLASH)
		}
	default:
		return ERR_UNEXPECTED_TOKEN(s.line, token)
	}

	return nil
}

func (s *Scanner) advance() string {
	s.current += 1
	return string(s.source[s.current-1])
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) addToken(tokType TokenType) {
	s.addTokenWithLiteral(tokType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokType TokenType, literal *Object) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(tokType, string(text), literal, s.line))
}
