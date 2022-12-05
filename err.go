package main

import (
	"fmt"
)

func reportError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
	hadError = true
}
