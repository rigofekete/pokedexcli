package main 

import (
		"strings"
		"fmt"
		"bufio"
		"os"
)

func cleanInput(text string) []string {
	text  = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: 		"help",
			description: 	"Displays a help message",
			callback: 	commandHelp,
		},
		"exit": {
			name: 		"exit",
			description:	"Exit the Pokedex",
			callback:	commandExit,
		},
	}
}

func startRepl() {

	reader := bufio.NewScanner(os.Stdin)
	for { 
		fmt.Print("\nPokedex > ")

		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]


		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

		//  k, _ := range commandMap {
		// 	if commandName == k {
		// 		found = true
		// 		commandMap[k].callback()
		// 	} 		
		// }
		//
		// if !found {
		// 	fmt.Print("Unknown command")
		// }


		// fmt.Printf("Your command was: %s", commandName)

	}
}

