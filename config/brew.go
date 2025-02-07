package config

import (
	"fmt"
	"os"
	"os/exec"
)

type brewConfig struct {
	Formulae []string `json:"formulae"`
	Casks    []string `json:"casks"`
}

func (cfg brewConfig) Install() error {
	if _, err := exec.LookPath("brew"); err != nil {
		return fmt.Errorf("brew not installed. install brew here (https://brew.sh) and try again.")
	}

	fmt.Println("brew: installing casks..")
	for _, cask := range cfg.Casks {
		// Ignore the "code" cask since we are already checking for it in CodeConfig.Install
		if cask == "code" {
			continue
		}

		if err := os.Setenv("HOMEBREW_NO_AUTO_UPDATE", "1"); err != nil {
			return fmt.Errorf("brew: failed to disable auto-update")
		}

		cmd := exec.Command("brew", "install", "--cask", cask)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install cask %q", cask)
		}
	}
	fmt.Println("brew: done, all casks installed")

	fmt.Println("brew: installing formulae..")
	for _, formula := range cfg.Formulae {
		cmd := exec.Command("brew", "install", formula)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install formula %q", formula)
		}
	}
	fmt.Println("brew: done, all formulae installed")

	return nil
}
