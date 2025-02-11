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

	if _, err := exec.LookPath("code"); err != nil {
		loadingIndicator.Suffix = "  code not installed. installing code via brew..."
		loadingIndicator.Start()

		cmd := exec.Command("brew", "install", "code")
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("brew: failed to install code")
		}

		loadingIndicator.Stop()
		fmt.Println("code: done, installed.")
	}

	loadingIndicator.Suffix = "  code: fetching extensions..."
	loadingIndicator.Restart()

	cmd := exec.Command("code", "--list-extensions")
	extensionsAlreadyInstalled, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("code: failed to list extensions")
	}

	loadingIndicator.Suffix = "  code: installing extensions..."
	loadingIndicator.Restart()

	for _, extension := range cfg.Extensions {
		if bytes.Contains(extensionsAlreadyInstalled, []byte(extension)) {
			continue
		}

		cmd := exec.Command("code", "--install-extension", extension)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("code: failed to install extension %q\n", extension)
		}
	}

	loadingIndicator.Stop()
	fmt.Println("code: done, all extensions installed")

	return nil
}
