package config

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func (cfg codeConfig) Install() error {
	loadingIndicator := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	defer loadingIndicator.Stop()

	loadingIndicator.Suffix = "  Checking for Visual Studio Code..."
	loadingIndicator.Start()

	if _, err := exec.LookPath("code"); err != nil {
		// Check if Visual Studio Code is installed first
		spotlightSearchForAppCmd := fmt.Sprintf("mdfind \"kMDItemKind == 'Application'\" | grep -i \"visual studio code\"")
		cmd := exec.Command("bash", "-c", spotlightSearchForAppCmd)
		spotlightResult, _ := cmd.Output()

		// An empty result means Visual Studio Code is not listed as an application
		if len(bytes.TrimSpace(spotlightResult)) == 0 {
			loadingIndicator.Suffix = "  Installing Visual Studio Code via Homebrew..."
			loadingIndicator.Restart()

			cmd := exec.Command("brew", "install", "--cask", "visual-studio-code")
			_, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("Failed to install Visual Studio Code")
			}
		} else {
			// A non-empty result means Visual Studio Code is installed but the code
			// command is not added to the PATH
			loadingIndicator.Suffix = "  Adding code command to PATH..."
			loadingIndicator.Restart()

			cmd := exec.Command("ln", "-s", "\"/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code\"", "/usr/local/bin/code")
			_, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("Failed to add code binary to PATH")
			}
		}
	}

	loadingIndicator.Suffix = "  Fetching extensions..."
	loadingIndicator.Restart()

	cmd := exec.Command("code", "--list-extensions")
	extensionsAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Failed to list extensions")
	}

	loadingIndicator.Suffix = "  Installing extensions..."
	loadingIndicator.Restart()

	for _, extension := range cfg.Extensions {
		if bytes.Contains(extensionsAlreadyInstalled, []byte(extension)) {
			continue
		}

		loadingIndicator.Suffix = fmt.Sprintf("  Installing extension %q...", extension)
		loadingIndicator.Restart()

		cmd := exec.Command("code", "--install-extension", extension)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("Failed to install extension %q\n", extension)
		}
	}

	loadingIndicator.Stop()
	fmt.Println("All extensions installed.")

	return nil
}
