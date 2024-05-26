package runner

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func RunRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("\n%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		err := run(line)
		if err != nil {

		}
	}
}
