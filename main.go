package main

import (
	"fmt"
	"gomonkey/repl"
	"os"
)

func main() {
	fmt.Printf("Monkey programming language REPL:\n")
	repl.Start(os.Stdin, os.Stdout)
}
