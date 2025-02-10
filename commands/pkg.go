// Package commands provides functions for working with the wit command line
// interface tool.
package commands

// Command is a command option in wit.
type Command struct {
	Name        string
	Description string
	// Run runs the command, returning any errors if any.
	Run func() error
	// ShowHelp prints a help message for the command.
	ShowHelp func()
}
