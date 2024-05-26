package main

import (
	"fmt"
	"os"

	"github.com/anukuljoshi/monkey-v2/runner"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(1)
	} else if len(args) == 2 {
		runner.RunFile(args[1])
	} else {
		fmt.Println("Welcome to Monkey Programming Language V2...")
		runner.RunRepl(os.Stdin, os.Stdout)
	}
}
