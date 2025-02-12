package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/digitalnest-wit/wit/commands"
)

func main() {
	flag.Parse()
	commandName := flag.Arg(0)

	if len(flag.Args()) < 1 {
		fmt.Print("wit is a tool used for managing computers at Digital NEST.\n\n")
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
