package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mac21/lox/scanner"
)

var (
	logger     = log.Default()
	sourceFile = ""
	hadError   = false
)

func init() {
	flag.StringVar(&sourceFile, "f", "", "Input source file with full path provided")
	flag.Parse()
}

func main() {
	if sourceFile != "" {
		runFile(sourceFile)
	} else {
		runPrompt(os.Stdin, os.Stdout)
	}
}

func runFile(source string) {
	fileBytes, err := ioutil.ReadFile(source)
	if err != nil {
		logger.Fatal("Failed to open input file")
	}

	run(string(fileBytes))

	if hadError {
		os.Exit(65)
	}
}

func run(source string) {
	s := scanner.New(source)
	tokens := s.scanTokens()

	for _, tok := range tokens {
		fmt.Println(tok)
	}
}
