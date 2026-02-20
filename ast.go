package main

import (
	"log"
	"math"
)

// made to enforce type for other nodes - acts as a tag
type Node interface {
	Eval() float64
}

type BinaryOpNode struct {
	left  Node
	op    Token
	right Node
}

func (b *BinaryOpNode) Eval() float64 {
	switch b.op.Type {
	case TokenPlus:
		return b.left.Eval() + b.right.Eval()
	case TokenMinus:
		return b.left.Eval() - b.right.Eval()
	case TokenMultiply:
		return b.left.Eval() * b.right.Eval()
	case TokenDivide:
		return b.left.Eval() / b.right.Eval()
	case TokenPow:
		return math.Pow(b.left.Eval(), b.right.Eval())
	default:
		log.Fatal("unknown token")
		return 0
	}
}

type NumberNode struct {
	value float64
}

func (n *NumberNode) Eval() float64 {
	return n.value
}
