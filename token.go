package main

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenLParen
	TokenRParen
	TokenPow
	TokenEOF
)

func (t TokenType) String() string {
	switch t {
	case TokenNumber:
		return "Number"
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
	case TokenPow:
		return "Pow"
	default:
		return "Unknown"
	}
}

type Token struct {
	Type  TokenType
	Value float64
}
