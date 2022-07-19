package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mac21/lox/repl"
	"github.com/mac21/lox/scanner"
)

var (
	logger     = log.Default()
	sourceFile string
)

func init() {
	flag.StringVar(&sourceFile, "f", "", "Input source file with full path provided")
}

func main() {
	flag.Parse()

	if sourceFile != "" {
		runFile(sourceFile)
	} else {
		repl.Run(os.StdIn, os.StdOut)
	}
}

func runFile(source string) {
	fileBytes, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logger.Fatal("Failed to open input file")
	}

	scn := scanner.New(string(fileBytes))
	tokens := scn.scanTokens()

	for _, tok := range tokens {
		fmt.Println(tok)
	}
}
