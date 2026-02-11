package main

type TokenType int

const (
	TokenInteger TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenLParen
	TokenRParen
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
	case TokenLParen:
		return "LParen"
	case TokenRParen:
		return "RParen"
	default:
		return "Unknown"
	}
}

type Token struct {
	Type  TokenType
	Value int
}
