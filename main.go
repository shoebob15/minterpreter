package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("minterpreter v0.1 | type nothing or ctrl+c to exit")

	log.Default().SetFlags(0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("minterpret> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		if line == "" {
			break
		}

		lexer, err := NewLexer(line)
		if err != nil {
			log.Default().Printf("%v", err)
			continue
		}

		parser, err := NewParser(lexer)
		if err != nil {
			log.Default().Printf("%v", err)
			continue
		}

		interpreter, err := NewInterpreter(parser)
		if err != nil {
			log.Default().Printf("%v", err)
			continue
		}

		result, err := interpreter.Interpret()
		if err == nil {
			fmt.Printf("=> %v\n", result)
		} else {
			fmt.Println("error: ", err)
			continue
		}
	}
}
