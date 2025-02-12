// Package config provides functions for working with the config command
// and managing configuration.
package config

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
