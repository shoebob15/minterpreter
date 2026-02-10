package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("minterpreter v0.1 | type nothing or ctrl+c to exit")

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

		interpreter, err := NewInterpreter(line)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}

		result, err := interpreter.Eval()
		if err == nil {
			fmt.Printf("=> %d\n", result)
		} else {
			fmt.Println("error: ", err)
			continue
		}
	}
}
