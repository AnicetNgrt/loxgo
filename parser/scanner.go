package parser

import (
	"strconv"
	"unicode"
)

type Scanner struct {
	source         []rune
	tokens         []Token
	start          int
	current        int
	line           int
	inBlockComment int
}

func NewScanner(source string) *Scanner {
	return &Scanner{[]rune(source), make([]Token, 0), 0, 0, 1, 0}
}

func (s *Scanner) ScanTokens() ([]Token, []Error) {
	errors := make([]Error, 0)

	for !s.isAtEnd() {
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
	tokenStr, tokenRune := s.advance()

	switch tokenStr {
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
		if s.match("/") {
			if s.inBlockComment > 0 {
				s.inBlockComment -= 1
			} else {
				return ERR_UNEXPECTED_TOKEN(s.line, tokenStr)
			}
		} else {
			s.addToken(TOK_STAR)
		}
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
			for s.peekStr() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else if s.match("*") {
			s.inBlockComment += 1
		} else {
			s.addToken(TOK_SLASH)
		}
	case "\"":
		err := s.string()
		if err != nil {
			return err
		}
	case " ", "\r", "\t":
		break
	case "\n":
		s.line += 1
	default:
		if s.inBlockComment > 0 {
			break
		} else if unicode.IsDigit(tokenRune) {
			s.number()
		} else if isAlpha(tokenRune) {
			s.identifier()
		} else {
			return ERR_UNEXPECTED_TOKEN(s.line, tokenStr)
		}
	}

	return nil
}

func (s *Scanner) identifier() {
	for isAlphaNum(s.peek()) {
		s.advance()
	}

	text := string(s.source[s.start:s.current])
	tType, ok := KeywordToTokenType[text]
	if !ok {
		tType = TOK_IDENTIFIER
	}
	s.addToken(tType)
}

func isAlpha(token rune) bool {
	return unicode.IsLetter(token) || string(token) == "_"
}

func isAlphaNum(token rune) bool {
	return isAlpha(token) || unicode.IsDigit(token)
}

func (s *Scanner) number() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peekStr() == "." && unicode.IsDigit(s.peekNext()) {
		s.advance()
		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	number, _ := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)
	s.addTokenWithLiteral(TOK_NUMBER, number)
}

func (s *Scanner) string() Error {
	for s.peekStr() != "\"" && !s.isAtEnd() {
		if s.peekStr() == "\n" {
			s.line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		return ERR_UNTERMINATED_STRING(s.line)
	}

	s.advance()

	s.addTokenWithLiteral(TOK_STRING, string(s.source[s.start+1:s.current-1]))
	return nil
}

func (s *Scanner) advance() (string, rune) {
	s.current += 1
	r := s.source[s.current-1]
	return string(r), r
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) peekStr() string {
	return string(s.peek())
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
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
	if s.inBlockComment > 0 {
		return
	}
	s.addTokenWithLiteral(tokType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(tokType, string(text), literal, s.line))
}
