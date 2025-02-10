package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/digitalnest-wit/wit/config"
)

var (
	// Available is a collection of wit commands. Entries are a command name
	// mapping to a Command.
	Available = map[string]Command{
		"config": {
			Name:        "config",
			Description: "execute commands defined in witconfig.json",
			Run: func() error {
				file, err := os.Open("witconfig.json")
				if err != nil {
					return fmt.Errorf("failed to read file: %w", err)
				}

				decoder := json.NewDecoder(file)
				config := config.Config{}

				if err := decoder.Decode(&config); err != nil {
					return fmt.Errorf("failed to decode witconfig.json: %w", err)
				}

				if err := config.Install(); err != nil {
					return err
				}

				return nil
			},
			ShowHelp: func() {
				fmt.Print("config\n\nexecutes commands defined in the witconfig.json file.\n\n")
				fmt.Print("If witconfig.json is not found in the current directory, the program ")
				fmt.Print("will exit with a status code of 1.\n\n")
				fmt.Println("Usage:\n    wit config")
				fmt.Println()
			},
		},
		"help": {
			Name:        "help",
			Description: "get help for a specific command",
			Run: func() error {
				// Intentionally unimplemented to avoid cyclic initializtion reference when fetching
				// all available commands. This function should never be called.
				return nil
			},
			ShowHelp: func() {
				// Intentionally unimplemented. This function should never be called.
			},
		},
	}
)
