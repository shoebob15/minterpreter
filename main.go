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

		interpreter := NewInterpreter(line)
		result, err := interpreter.Eval()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(result)
	}
}
