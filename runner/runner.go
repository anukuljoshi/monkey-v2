package runner

import (
	"fmt"

	"github.com/anukuljoshi/monkey-v2/lexer"
)

func run(input string) {
	lexer := lexer.New(input)

	tokens, err := lexer.ScanTokens()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	return
}
