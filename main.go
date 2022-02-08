package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

func run(source string) []Error {
	scanner := NewScanner(source)
	tokens, errs := scanner.scanTokens()
	if len(errs) > 0 {
		return errs
	}

	for idx, value := range tokens {
		fmt.Printf("%d: %s\n", idx, value.toString())
	}

	return nil
}

func printErrors(errs ...Error) {
	report("", errs...)
}

func report(where string, errs ...Error) {
	for _, err := range errs {
		fmt.Fprintf(os.Stderr, "[line %d] Error %s:\n	%s\n", err.line, where, err.message)
	}
}