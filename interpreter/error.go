package interpreter

import "fmt"

type ErrorStruct struct {
	Line    int
	Message string
}

type Error *ErrorStruct

func ERR(line int, msg string) Error {
	return &ErrorStruct{line, msg}
}

func ERR_UNEXPECTED_TOKEN(line int, token string) Error {
	return &ErrorStruct{line, fmt.Sprintf("Unexpected token: %s", token)}
}

func ERR_UNTERMINATED_STRING(line int) Error {
	return &ErrorStruct{line, "Unterminated string literal"}
}
