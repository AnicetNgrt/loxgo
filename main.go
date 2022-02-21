package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"loxgo/interpreter"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("could not open script")
		return
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("could not read script")
		return
	}

	errs := run(string(bytes))

	if errs != nil {
		printErrors(errs...)
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		errs := run(line)
		printErrors(errs...)
	}
}

func run(source string) []interpreter.Error {
	scanner := interpreter.NewScanner(source)
	tokens, errs := scanner.ScanTokens()
	if len(errs) > 0 {
		return errs
	}
	parser := interpreter.NewParser(tokens)
	expr, errs := parser.Parse()
	if len(errs) > 0 {
		return errs
	}

	for idx, value := range tokens {
		fmt.Printf("%d: %s\n", idx, value.ToString())
	}

	fmt.Printf("%s\n", expr.PrettyPrint())

	return nil
}

func printErrors(errs ...interpreter.Error) {
	report("", errs...)
}

func report(where string, errs ...interpreter.Error) {
	for _, err := range errs {
		fmt.Fprintf(os.Stderr, "[line %d] Error %s:\n	%s\n", err.Line, where, err.Message)
	}
}
