package main

type TokenType int

const (
	TokenInteger TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenEOF
)

func (t TokenType) String() string {
	switch t {
	case TokenInteger:
		return "Integer"
	case TokenPlus:
		return "Plus"
	case TokenMinus:
		return "Minus"
	case TokenMultiply:
		return "Multiply"
	case TokenDivide:
		return "Divide"
	case TokenEOF:
		return "EOF"
	default:
		return "Unknown"
	}
}

type Token struct {
	Type  TokenType
	Value int
}
