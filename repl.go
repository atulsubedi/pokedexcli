package main

import (
	"bufio"
	"fmt"
	"os"
)

var cliName string = "pokedex"

func printPromt() {
	fmt.Println(cliName, " > ")
}

func displayHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Display a help message")
	fmt.Println("exit: Exit the Pokedex")
}

func main() {
	commands := map[string]interface{}{
		"help": displayHelp,
	}
	reader := bufio.NewScanner(os.Stdin)
	printPromt()
	for reader.Scan() {
		text := reader.Text()
		if command, exist := commands[text]; exist {
			if cmdFunc, ok := command.(func()); ok {
				cmdFunc()
			} else {
				fmt.Println("Invalid command")
			}
		}
		if text == "exit" {
			os.Exit(0)
		}
	}
}
