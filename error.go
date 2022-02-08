package main

import "fmt"

type ErrorStruct struct {
	line    int
	message string
}

type Error *ErrorStruct

func ERR_UNEXPECTED_TOKEN(line int, token string) Error {
	return &ErrorStruct{line, fmt.Sprintf("Unexpected token: %s", token)}
}
