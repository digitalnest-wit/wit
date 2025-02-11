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
	// Run runs the command, returning any errors if any.
	Run func() error
	// ShowHelp prints a help message for the command.
	ShowHelp func()
}

// ShowHelp prints the usage of the help command.
func ShowHelp() {
	commandName := flag.Arg(1)

	if len(flag.Args()) < 2 || commandName == "help" {
		fmt.Print("help\n\nprovides detailed help for a specified command.\n\n")
		fmt.Println("Usage:\n    wit help <command>")
		fmt.Println()
		ShowAllCommands()

		return
	}

	command, isValidCommand := Available[commandName]
	if !isValidCommand {
		UnknownCommand(commandName)
		os.Exit(1)
	}

	command.ShowHelp()
}

// UnknownCommand prints a message to the user notifying them that the command
// received was unrecognized and invalid. A call to UnknownCommand will never
// panic.
func UnknownCommand(cmd string) {
	fmt.Printf("Unrecognized command received: %q.\n", cmd)
}

// ShowAllCommands prints a list of all the available commands.
func ShowAllCommands() {
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
