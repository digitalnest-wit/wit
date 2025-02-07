package config

import (
	"bytes"
	"fmt"
	"os/exec"
)

type codeConfig struct {
	Extensions []string `json:"extensions"`
}

func (cfg codeConfig) Install() error {
	if _, err := exec.LookPath("code"); err != nil {
		fmt.Println("code not installed. installing code via brew..")

		cmd := exec.Command("brew", "install", "code")
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install code")
		}

		fmt.Println("done, code installed.")
	}

	cmd := exec.Command("code", "--list-extensions")
	extensions, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("code: failed to list extensions")
	}

	fmt.Println("code: installing extensions..")
	for _, extension := range cfg.Extensions {
		if bytes.Contains(extensions, []byte(extension)) {
			continue
		}

		cmd := exec.Command("code", "--install-extension", extension)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("code: failed to install extension %q\n", extension)
		}
	}
	fmt.Println("code: done, extensions installed")

	return nil
}
