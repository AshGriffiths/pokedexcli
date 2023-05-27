package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startPokedex() {
	prompt := "pokedex > "
	inputReader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		inputReader.Scan()

		input := strings.Fields(strings.ToLower(strings.TrimSpace(inputReader.Text())))

		if len(input) == 0 {
			continue
		}

		actionName := input[0]
		do, ok := getActions()[actionName]

		if ok {
			err := do.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")

	}
}
