package main

// TODO: the entire workflow of creating new lexer, parser, and interpreter
// instances is wildly inefficient. rather, it should be similar to
// reconstructing tokens and the tree based on new input
type Interpreter struct {
	parser *Parser
}

func (i *Interpreter) Interpret() (float64, error) {
	tree, err := i.parser.Eval()
	if err != nil {
		return 0, err
	}

	return tree.Eval(), nil
}

func NewInterpreter(parser *Parser) (*Interpreter, error) {
	return &Interpreter{
		parser: parser,
	}, nil
}
