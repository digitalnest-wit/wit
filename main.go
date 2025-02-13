package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/digitalnest-wit/wit/commands"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("Unexpected operating system. Sorry, wit is a macOS specific tool.")
		os.Exit(1)
	}

	flag.Parse()
	commandName := flag.Arg(0)

	if len(flag.Args()) < 1 {
		fmt.Println("          _  __ ")
		fmt.Println(" _    __ (_)/ /_")
		fmt.Println("| |/|/ // // __/")
		fmt.Println("|__,__//_/ \\__/ ")
		fmt.Println()
		fmt.Print("A powerful command-line interface (CLI) tool designed ")
		fmt.Print("specifically for Digital NEST to streamline and automate common ")
		fmt.Print("computer operations.\n\nThis utility aims to simplify IT ")
		fmt.Print("management and enhance productivity for the organization's ")
		fmt.Print("technical staff.\n\n")
		fmt.Print("Usage:\n    wit <command> [arguments]\n\n")
		commands.PrintAvailable()

		return
	}

	if commandName == "help" {
		commands.PrintHelp()
		return
	}

	command, isValidCommand := commands.Available[commandName]
	if !isValidCommand {
		commands.HandleUnknownCmd(commandName)
	}

	if err := command.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
