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

type Token struct {
	Type  TokenType
	Value int
}
