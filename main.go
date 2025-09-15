package main

import (
		"fmt"
		"bufio"
		"os"
)

func main() {
	for { 
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nPokedex > ")

		var input string

		scanner.Scan()
		input = scanner.Text()

		final_input := cleanInput(input)

		fmt.Printf("Your command was: %s", final_input[0])
	}
}








	
