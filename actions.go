package main

import (
	"errors"
	"fmt"
	"os"
)

type action struct {
	name        string
	description string
	callback    func(*config) error
}

func actionHelp(cfg *config) error {
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

func actionMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func actionMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func actionExit(cfg *config) error {
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
		"map": {
			name:        "map",
			description: "Get the next 20 locations",
			callback:    actionMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    actionMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    actionExit,
		},
	}
}
