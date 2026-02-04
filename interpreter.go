package main

import (
	"fmt"
)

type Interpreter struct {
	lexer        *Lexer
	currentToken Token
}

func (i *Interpreter) Eval() (int, error) {
	var err error
	i.currentToken, err = i.lexer.getNextToken()

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

func (i *Interpreter) eat(t TokenType) error {
	if i.currentToken.Type == t {
		token, err := i.lexer.getNextToken()
		if err != nil {
			return err
		}

		i.currentToken = token
	} else {
		return fmt.Errorf("unexpected token at 0:%d", i.lexer.pos)
	}

	return nil
}

func NewInterpreter(input string) *Interpreter {
	return &Interpreter{lexer: NewLexer(input)}
}
