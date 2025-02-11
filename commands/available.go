package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/digitalnest-wit/wit/config"
)

var (
	// filePathNameRegex is a regular expression for validating configuration file
	// names.
	filePathNameRegex = regexp.MustCompile(".*\\.witconfig.json")

	// Available is a collection of wit commands. Entries are a command name
	// mapping to a Command.
	Available = map[string]Command{
		"config": {
			Name:        "config",
			Description: "execute commands defined in a configuration file",
			Run: func() error {
				filePath := strings.TrimSpace(flag.Arg(1))
				// Set the file path to a relative file named witconfig.json if a file
				// path is not supplied as an argument.
				if len(filePath) == 0 {
					filePath = "witconfig.json"
				}

				if !filePathNameRegex.MatchString(filePath) {
					return fmt.Errorf("invalid file name %q: file name must end in .witconfig.json.", filePath)
				}

				file, err := os.Open(filePath)
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
				fmt.Print("config\n\nexecutes commands defined in the witconfig.json ")
				fmt.Print("file.\n\nIf a path argument is supplied, then the program ")
				fmt.Print(" looks for the configuration file there. Otherwise the ")
				fmt.Print(" program looks in the active directory.\n\nIf a configuration ")
				fmt.Print("file cannot be resolved, then the program will exit with a ")
				fmt.Print("status code of 1.\n\n")
				fmt.Println("Usage:\n    wit config [path]")
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
