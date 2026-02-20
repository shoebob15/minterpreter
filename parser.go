package main

import (
	"fmt"
)

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func (p *Parser) eat(t TokenType) error {
	if p.currentToken.Type == t {
		token, err := p.lexer.getNextToken()
		if err != nil {
			return err
		}

		p.currentToken = token
	} else {
		return fmt.Errorf("expected %v, got %v at index %d",
			t, p.currentToken.Type, p.lexer.pos)
	}

	return nil
}

func (p *Parser) Eval() (Node, error) {
	node, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.currentToken.Type == TokenPlus || p.currentToken.Type == TokenMinus {
		token := p.currentToken

		if err := p.eat(token.Type); err != nil {
			return nil, err
		}

		right, err := p.term()
		if err != nil {
			return nil, err
		}

		node = &BinaryOpNode{left: node, op: token, right: right}
	}

	return node, nil
}

func (p *Parser) term() (Node, error) {
	node, err := p.power()
	if err != nil {
		return nil, err
	}

	for p.currentToken.Type == TokenMultiply || p.currentToken.Type == TokenDivide {
		token := p.currentToken
		if token.Type == TokenMultiply {
			if err := p.eat(TokenMultiply); err != nil {
				return nil, err
			}
		} else if token.Type == TokenDivide {
			if err := p.eat(TokenDivide); err != nil {
				return nil, err
			}
		}

		right, err := p.power()
		if err != nil {
			return nil, err
		}

		node = &BinaryOpNode{left: node, op: token, right: right}
	}

	return node, nil
}

func (p *Parser) power() (Node, error) {
	node, err := p.factor()
	if err != nil {
		return nil, err
	}

	if p.currentToken.Type == TokenPow {
		token := p.currentToken
		if err := p.eat(TokenPow); err != nil {
			return nil, err
		}

		right, err := p.power()
		if err != nil {
			return nil, err
		}
		node = &BinaryOpNode{left: node, op: token, right: right}
	}

	return node, nil
}

func (p *Parser) factor() (Node, error) {
	token := p.currentToken

	if token.Type == TokenNumber {
		if err := p.eat(TokenNumber); err != nil {
			return nil, err
		}
		return &NumberNode{value: token.Value}, nil
	}

	if token.Type == TokenLParen {
		if err := p.eat(TokenLParen); err != nil {
			return nil, err
		}
		node, err := p.Eval()
		if err != nil {
			return nil, err
		}
		if err := p.eat(TokenRParen); err != nil {
			return nil, err
		}
		return node, nil
	}

	return &NumberNode{}, fmt.Errorf("incorrect parenthesis")
}

func NewParser(lexer *Lexer) (*Parser, error) {
	token, err := lexer.getNextToken()
	if err != nil {
		return nil, err
	}

	return &Parser{
		lexer:        lexer,
		currentToken: token,
	}, nil
}
