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
		return fmt.Errorf("Homebrew not installed. Install it here (https://brew.sh) and try again")
	}

	// Disable Homebrew's auto update feature (on by default)
	if err := os.Setenv("HOMEBREW_NO_AUTO_UPDATE", "1"); err != nil {
		return fmt.Errorf("Failed to disable Homebrew auto-update")
	}

	loadingIndicator := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	defer loadingIndicator.Stop()

	loadingIndicator.Suffix = "  Fetching casks..."
	loadingIndicator.Start()

	cmd := exec.Command("brew", "list", "-1", "--cask")
	casksAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Failed to list casks")
	}

	loadingIndicator.Suffix = "  Installing casks..."
	loadingIndicator.Restart()

	for _, cask := range cfg.Casks {
		// Skip the cask installation if it's "visual-studio-code" or was already
		// installed via Homebrew.
		if cask == "visual-studio-code" || bytes.Contains(casksAlreadyInstalled, []byte(cask)) {
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

		loadingIndicator.Suffix = fmt.Sprintf("  Installing cask %q...", cask)
		loadingIndicator.Restart()

		cmd := exec.Command("brew", "install", "--cask", cask)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("Failed to install cask %q", cask)
		}
	}

	loadingIndicator.Suffix = "  Fetching formulae..."
	loadingIndicator.Restart()

	cmd = exec.Command("brew", "list", "-1", "--formulae")
	formulaeAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Failed to list formulae")
	}

	loadingIndicator.Suffix = "  Installing formulae..."
	loadingIndicator.Restart()

	for _, formula := range cfg.Formulae {
		if bytes.Contains(formulaeAlreadyInstalled, []byte(formula)) {
			continue
		}

		loadingIndicator.Suffix = fmt.Sprintf("  Installing formula %q...", formula)
		loadingIndicator.Restart()

		cmd := exec.Command("brew", "install", formula)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("Failed to install Homebrew formula %q", formula)
		}
	}

	loadingIndicator.Stop()
	fmt.Println("All casks and formulae installed.")

	return nil
}
