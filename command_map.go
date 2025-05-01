package main

import (
	"errors"
	"fmt"
)


func commandMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil

}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil

}

