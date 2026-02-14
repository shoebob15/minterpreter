package main

import (
	"fmt"
	"math"
)

type Interpreter struct {
	lexer        *Lexer
	currentToken Token
}

// Eval : Term (Plus | Minus) Term
func (i *Interpreter) Eval() (float64, error) {
	result, err := i.term()

	if err != nil {
		return 0, err
	}

	for i.currentToken.Type == TokenPlus || i.currentToken.Type == TokenMinus {
		token := i.currentToken
		i.eat(token.Type)

		term, err := i.term()
		if err != nil {
			return 0, err
		}

		switch token.Type {
		case TokenPlus:
			result += term

		case TokenMinus:
			result -= term
		}
	}

	return result, nil
}

// TERM : Factor (Multiply | Divide) Factor
func (i *Interpreter) term() (float64, error) {
	result, err := i.power()
	if err != nil {
		return 0, err
	}

	for i.currentToken.Type == TokenMultiply || i.currentToken.Type == TokenDivide {
		token := i.currentToken
		i.eat(token.Type)

		factor, err := i.power()
		if err != nil {
			return 0, err
		}

		switch token.Type {
		case TokenMultiply:
			result *= factor
		case TokenDivide:
			result /= factor
		}
	}

	return result, nil
}

// POWER : factor Pow factor
func (i *Interpreter) power() (float64, error) {
	result, err := i.factor()
	if err != nil {
		return 0, err
	}

	if i.currentToken.Type == TokenPow {
		if err := i.eat(TokenPow); err != nil {
			return 0, err
		}

		power, err := i.power()
		if err != nil {
			return 0, err
		}
		result = math.Pow(result, power)
	}

	return result, nil
}

// FACTOR : Number | LParen Expr RParen
func (i *Interpreter) factor() (float64, error) {
	token := i.currentToken

	if token.Type == TokenNumber {
		if err := i.eat(TokenNumber); err != nil {
			return 0, err
		}
		return token.Value, nil
	}

	if token.Type == TokenLParen {
		i.eat(TokenLParen)
		result, err := i.Eval()
		if err != nil {
			return 0, err
		}
		i.eat(TokenRParen)
		return result, nil
	}

	return 0, fmt.Errorf("incorrect parenthesis")

}

func (i *Interpreter) eat(t TokenType) error {
	if i.currentToken.Type == t {
		token, err := i.lexer.getNextToken()
		if err != nil {
			return err
		}

		i.currentToken = token
	} else {
		return fmt.Errorf("expected %v, got %v at index %d",
			t, i.currentToken.Type, i.lexer.pos)
	}

	return nil
}

func NewInterpreter(input string) (*Interpreter, error) {
	lexer, err := NewLexer(input)
	if err != nil {
		return nil, err
	}

	token, err := lexer.getNextToken()
	if err != nil {
		return nil, err
	}

	return &Interpreter{
		lexer:        lexer,
		currentToken: token,
	}, nil
}
