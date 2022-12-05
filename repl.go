package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func runPrompt(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		err := scanner.Err()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		line := scanner.Text()
		run(line)
	}
}
