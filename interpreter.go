package main

type Interpreter struct {
	input        string
	pos          int
	currentToken Token
}

func NewInterpreter(input string) *Interpreter {
	return &Interpreter{input: input}
}

func (i *Interpreter) Eval() int {
	i.currentToken = i.getNextToken()

	left := i.currentToken
	i.eat(TokenInteger)

	_ = i.currentToken // operator
	i.eat(TokenPlus)

	right := i.currentToken
	i.eat(TokenInteger)

	return left.Value + right.Value
}

func (i *Interpreter) getNextToken() Token {
	if i.pos > len(i.input)-1 {
		return Token{Type: TokenEOF}
	}

	ch := i.input[i.pos]

	if ch >= '0' && ch <= '9' {
		i.pos++
		return Token{
			Type:  TokenInteger,
			Value: int(ch - '0'),
		}
	}

	if ch == '+' {
		i.pos += 1
		return Token{Type: TokenPlus}
	}

	panic("unrecognized token in input")
}

func (i *Interpreter) eat(t TokenType) {
	if i.currentToken.Type == t {
		i.currentToken = i.getNextToken()
	} else {
		panic("unexpected token in input")
	}
}
