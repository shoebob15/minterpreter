package main

type TokenType int

const (
	TokenInteger TokenType = iota
	TokenPlus
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value int
}
