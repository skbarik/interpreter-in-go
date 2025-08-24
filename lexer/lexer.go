package lexer

import "github.com/skbarik/interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l Lexer) NextToken() token.Token {
	// TODO
	return token.Token{}
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}
