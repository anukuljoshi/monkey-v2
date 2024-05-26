package runner

import (
	"fmt"
	"os"
)

func RunFile(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	input := string(b)
	run(input)
}
