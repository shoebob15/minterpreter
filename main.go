package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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

		interpreter := NewInterpreter(line)
		fmt.Println(interpreter.Eval())
	}
}
