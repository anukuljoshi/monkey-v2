package lexer

import "github.com/anukuljoshi/monkey-v2/token"

type Lexer struct {
    input    string
    position int
    readPosition int
    ch byte
}

func New(input string) *Lexer {
    return &Lexer{
	input: input,
    }
}

func (l *Lexer) ScanTokens() ([]token.Token, error) {
    return []token.Token{}, nil
}
