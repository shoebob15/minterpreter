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
		handleError(err)

		result, err := interpreter.Eval()
		handleError(err)
		fmt.Println(result)
	}
}

// will print error but not end program
func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
