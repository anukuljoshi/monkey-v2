package errors

import (
	"fmt"
)


type ErrorMessage string

type LexerError struct {
	Line int
	Message ErrorMessage
}

func NewLexerError(line int, message string) *LexerError {
	return &LexerError{
		Line: line,
		Message: ErrorMessage(message),
	}
}

func (l LexerError) Error() {
	fmt.Printf("Line: [%d], LexerError: %s", l.Line, l.Message)
}
