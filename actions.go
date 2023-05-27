package main

import (
	"fmt"
	"os"
)

type action struct {
	name        string
	description string
	callback    func() error
}

func actionHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getActions() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func actionExit() error {
	os.Exit(0)
	return nil
}

func getActions() map[string]action {
	return map[string]action{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    actionHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    actionExit,
		},
	}
}
