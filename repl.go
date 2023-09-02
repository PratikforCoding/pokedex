package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		command, exist := getCommands()[cmd]
		if exist {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Printf("Unknown command: %s", command)
			continue
		}
	}
}

func cleanInput(cmd string) []string {
	output := strings.ToLower(cmd)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" : {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit" : {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
	}
}
