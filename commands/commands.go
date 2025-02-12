// Package commands provides functions for working with the wit command line
// interface tool.
package commands

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
)

// Command is a command option in wit.
type Command struct {
	Name        string
	Description string
	Run         func() error
	PrintHelp   func()
}

// PrintHelp prints the usage of the help command and calls the PrintHelp
// method for the command received via command line arguments.
func PrintHelp() {
	commandName := flag.Arg(1)

	if len(flag.Args()) < 2 || commandName == "help" {
		fmt.Print("help\n\nprovides detailed help for a specified command.\n\n")
		fmt.Println("Usage:\n    wit help <command>")
		fmt.Println()
		PrintAvailable()

		return
	}

	command, isValidCommand := Available[commandName]
	if !isValidCommand {
		HandleUnknownCmd(commandName)
		os.Exit(1)
	}

	command.PrintHelp()
}

// HandleUnknownCmd prints a message notifying the user that the command
// received was unrecognized. The program will exit with a status code
// of 1.
func HandleUnknownCmd(cmd string) {
	fmt.Printf("Unrecognized command received: %q.\n", cmd)
	os.Exit(1)
}

// PrintAvailable prints a list of all the available commands.
func PrintAvailable() {
	fmt.Println("The commands are:")
	fmt.Println()

	// Sort the available command names lexicographically
	availableCommands := maps.Keys(Available)
	availableCommandsSlice := slices.Sorted(availableCommands)

	for _, commandName := range availableCommandsSlice {
		commandDescription := Available[commandName].Description
		fmt.Printf("    %-10s %s\n", commandName, commandDescription)
	}

	fmt.Println()
}
