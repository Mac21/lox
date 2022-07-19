package repl

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Run(in io.Reader, out io.Writer) {
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
		l := lexer.New(line)
		l.ScanToStdout()
	}
}
