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

func (l *Lexer) number() (float64, error) {
	builder := strings.Builder{}
	dotCount := 0

	for l.currentChar != 0 && (isDigit(l.currentChar) || l.currentChar == '.') {
		if l.currentChar == '.' {
			dotCount++
			if dotCount > 1 {
				return 0, fmt.Errorf("too many dots")
			}
		}
		builder.WriteByte(l.currentChar)
		l.advance()
	}

	result, err := strconv.ParseFloat(builder.String(), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number at 0:%d", l.pos)
	}

	return result, nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) getNextToken() (Token, error) {
	for l.currentChar != 0 {
		if l.currentChar == ' ' {
			l.skipWhitespace()
			continue
		}

		if l.currentChar >= '0' && l.currentChar <= '9' {
			value, err := l.number()
			if err != nil {
				return Token{}, err
			}

			return Token{
				Type:  TokenNumber,
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

		if l.currentChar == '(' {
			l.advance()
			return Token{Type: TokenLParen}, nil
		}

		if l.currentChar == ')' {
			l.advance()
			return Token{Type: TokenRParen}, nil
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
