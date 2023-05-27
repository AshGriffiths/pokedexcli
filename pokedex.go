package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ashgriffiths/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startPokedex(cfg *config) {
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
			err := do.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")

	}
}
