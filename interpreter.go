package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Interpreter struct {
	input        string
	pos          int
	currentToken Token
	currentChar  byte // nil when == 0 (null byte)
}

func (i *Interpreter) Eval() (int, error) {
	var err error
	i.currentToken, err = i.getNextToken()

	if err != nil {
		return 0, err
	}

	left := i.currentToken
	err = i.eat(TokenInteger)

	op := i.currentToken

	if op.Type == TokenPlus {
		err = i.eat(TokenPlus)
	} else {
		err = i.eat(TokenMinus)
	}

	right := i.currentToken
	err = i.eat(TokenInteger)

	if err != nil {
		return 0, err
	}

	if op.Type == TokenPlus {
		return left.Value + right.Value, nil
	}
	return left.Value - right.Value, nil
}

func (i *Interpreter) getNextToken() (Token, error) {
	for i.currentChar != 0 {
		if i.currentChar == ' ' {
			i.skipWhitespace()
			continue
		}

		if i.currentChar >= '0' && i.currentChar <= '9' {
			return Token{
				Type:  TokenInteger,
				Value: i.integer(),
			}, nil
		}

		if i.currentChar == '+' {
			i.advance()
			return Token{Type: TokenPlus}, nil
		}

		if i.currentChar == '-' {
			i.advance()
			return Token{Type: TokenMinus}, nil
		}

		return Token{}, fmt.Errorf("unrecognized token at 0:%d", i.pos)
	}

	return Token{Type: TokenEOF}, nil

}

func (i *Interpreter) advance() {
	i.pos += 1

	if i.pos > len(i.input)-1 {
		i.currentChar = 0
	} else {
		i.currentChar = i.input[i.pos]
	}
}

func (i *Interpreter) eat(t TokenType) error {
	if i.currentToken.Type == t {
		token, err := i.getNextToken()
		if err != nil {
			return err
		}

		i.currentToken = token
	} else {
		return fmt.Errorf("unexpected token at 0:%d", i.pos)
	}

	return nil
}

func (i *Interpreter) integer() int {
	builder := strings.Builder{}
	ch := i.currentChar

	for i.currentChar != 0 && i.currentChar >= '0' && i.currentChar <= '9' {
		builder.WriteByte(ch)
		i.advance()
		ch = i.currentChar
	}

	result, err := strconv.Atoi(builder.String())
	if err != nil {
		panic("an internal error occurred")
	}

	return result
}

func (i *Interpreter) skipWhitespace() {
	for i.currentChar != 0 && i.currentChar == ' ' {
		i.advance()
	}
}

func NewInterpreter(input string) *Interpreter {
	return &Interpreter{input: input, currentChar: input[0]}
}
