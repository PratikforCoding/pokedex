package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("exiting pokedex!")
	os.Exit(0)
	return nil
}