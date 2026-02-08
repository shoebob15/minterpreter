package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Lexer struct {
	input       string
	pos         int
	currentChar byte // nil when == 0 (null byte)
}

func (l *Lexer) advance() {
	l.pos++

	if l.pos > len(l.input)-1 {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.pos]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar != 0 && l.currentChar == ' ' {
		l.advance()
	}
}

func (l *Lexer) integer() (int, error) {
	builder := strings.Builder{}

	for l.currentChar != 0 && l.currentChar >= '0' && l.currentChar <= '9' {
		builder.WriteByte(l.currentChar)
		l.advance()
	}

	result, err := strconv.Atoi(builder.String())
	if err != nil {
		return 0, fmt.Errorf("invalid integer at 0:%d", l.pos)
	}

	return result, nil
}

func (l *Lexer) getNextToken() (Token, error) {
	for l.currentChar != 0 {
		if l.currentChar == ' ' {
			l.skipWhitespace()
			continue
		}

		if l.currentChar >= '0' && l.currentChar <= '9' {
			value, err := l.integer()
			if err != nil {
				return Token{}, err
			}

			return Token{
				Type:  TokenInteger,
				Value: value,
			}, nil
		}

		if l.currentChar == '+' {
			l.advance()
			return Token{Type: TokenPlus}, nil
		}

		if l.currentChar == '-' {
			l.advance()
			return Token{Type: TokenMinus}, nil
		}

		if l.currentChar == '*' {
			l.advance()
			return Token{Type: TokenMultiply}, nil
		}

		if l.currentChar == '/' {
			l.advance()
			return Token{Type: TokenDivide}, nil
		}

		return Token{}, fmt.Errorf("unrecognized token at 0:%d", l.pos)
	}

	return Token{Type: TokenEOF}, nil
}

func NewLexer(input string) (*Lexer, error) {
	if len(input) < 1 {
		return &Lexer{}, fmt.Errorf("invalid input for lexer: %s", input)
	}
	return &Lexer{input: input, pos: 0, currentChar: input[0]}, nil
}
