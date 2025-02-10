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
		fmt.Println("error: unexpected runtime operating system")
		os.Exit(1)
	}

	flag.Parse()
	commandName := flag.Arg(0)

	if len(flag.Args()) < 1 {
		fmt.Print("wit is a tool used for managing computers at Digital NEST.\n\n")
		fmt.Print("Usage:\n    wit <command> [arguments]\n\n")
		commands.ShowAllCommands()

		return
	}

	if commandName == "help" {
		commands.ShowHelp()
		return
	}

	command, isValidCommand := commands.Available[commandName]
	if !isValidCommand {
		commands.UnknownCommand(commandName)
		os.Exit(1)
	}

	if err := command.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
