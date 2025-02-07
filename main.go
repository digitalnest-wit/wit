package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/digitalnest-wit/wit/config"
)

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Println("error: unexpected runtime operating system")
		os.Exit(1)
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		PrintUsage()
		return
	}

	switch flag.Arg(0) {
	case "config":
		if err := Config(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		PrintUnknownCommand()
		os.Exit(1)
	}
}

func PrintUsage() {
	fmt.Println("usage")
}

func PrintUnknownCommand() {
	fmt.Println("unknown")
}

func Config() error {
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
}
