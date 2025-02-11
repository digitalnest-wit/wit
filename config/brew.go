package config

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func (cfg brewConfig) Install() error {
	if _, err := exec.LookPath("brew"); err != nil {
		return fmt.Errorf("brew not installed. install brew here (https://brew.sh) and try again.")
	}

	loadingIndicator := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	loadingIndicator.Suffix = "  brew: installing casks..."
	loadingIndicator.Start()

	defer loadingIndicator.Stop()

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

	loadingIndicator.Suffix = "  brew: installing formulae.."
	loadingIndicator.Restart()

	for _, formula := range cfg.Formulae {
		cmd := exec.Command("brew", "install", formula)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install formula %q", formula)
		}
	}

	loadingIndicator.Stop()
	fmt.Println("brew: done, all casks and formulae installed")

	return nil
}
