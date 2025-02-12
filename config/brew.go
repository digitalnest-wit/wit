package config

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func (cfg brewConfig) Install() error {
	if _, err := exec.LookPath("brew"); err != nil {
		return fmt.Errorf("brew not installed. install brew here (https://brew.sh) and try again.")
	}

	// Disable Homebrew's auto update feature (on by default)
	if err := os.Setenv("HOMEBREW_NO_AUTO_UPDATE", "1"); err != nil {
		return fmt.Errorf("brew: failed to disable auto-update")
	}

	loadingIndicator := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	loadingIndicator.Suffix = "  brew: fetching casks..."
	loadingIndicator.Start()

	defer loadingIndicator.Stop()

	cmd := exec.Command("brew", "list", "-1", "--cask")
	casksAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("brew: failed to list casks")
	}

	loadingIndicator.Suffix = "  brew: installing casks..."
	loadingIndicator.Restart()

	for _, cask := range cfg.Casks {
		// Skip the cask installation if it's "code" or was already installed via
		// Homebrew.
		if cask == "code" || bytes.Contains(casksAlreadyInstalled, []byte(cask)) {
			continue
		}

		if runtime.GOOS == "darwin" {
			applicationName := strings.ReplaceAll(cask, "-", " ")
			spotlightSearchForAppCmd := fmt.Sprintf("mdfind \"kMDItemKind == 'Application'\" | grep -i \"%s\"", applicationName)

			cmd := exec.Command("bash", "-c", spotlightSearchForAppCmd)
			spotlightResult, _ := cmd.Output()

			// Skip the cask if an application already exists on the system with the
			// same name as the cask
			if len(bytes.TrimSpace(spotlightResult)) != 0 {
				continue
			}
		}

		loadingIndicator.Suffix = fmt.Sprintf("  brew: installing %q...", cask)
		loadingIndicator.Restart()

		cmd := exec.Command("brew", "install", "--cask", cask)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install cask %q", cask)
		}
	}

	loadingIndicator.Suffix = "  brew: fetching formulae..."
	loadingIndicator.Restart()

	cmd = exec.Command("brew", "list", "-1", "--formulae")
	formulaeAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("brew: failed to list formulae")
	}

	loadingIndicator.Suffix = "  brew: installing formulae.."
	loadingIndicator.Restart()

	for _, formula := range cfg.Formulae {
		if bytes.Contains(formulaeAlreadyInstalled, []byte(formula)) {
			continue
		}

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
