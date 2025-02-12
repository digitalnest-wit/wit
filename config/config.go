// Package config provides functions for working with the config command
// and managing configuration.
package config

// Config contains all the configuration data defined in a .witconfig
// file.
type Config struct {
	Version string     `yaml:"version"`
	Brew    brewConfig `yaml:"brew"`
	Code    codeConfig `yaml:"code"`
}

// Installer is an interface that exposes a basic Install function.
type Installer interface {
	Install() error
}

type brewConfig struct {
	Formulae []string `yaml:"formulae"`
	Casks    []string `yaml:"casks"`
}

type codeConfig struct {
	Extensions []string `yaml:"extensions"`
}

// Install everything defined in the configuration file.
func (c Config) Install() error {
	if err := c.Brew.Install(); err != nil {
		return err
	}

	if err := c.Code.Install(); err != nil {
		return err
	}

	return nil
}
