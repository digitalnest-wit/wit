package config

type witConfig struct {
	Version    string     `json:"version"`
	BrewConfig brewConfig `json:"brew_install"`
	CodeConfig codeConfig `json:"code_install"`
}

// Config contains all the configuration data defined in witconfig.json
type Config struct {
	config witConfig
}

// Install everything defined in witconfig.json
func (c Config) Install() error {
	if err := c.config.BrewConfig.Install(); err != nil {
		return err
	}

	if err := c.config.CodeConfig.Install(); err != nil {
		return err
	}

	return nil
}

// Version returns the version of the witconfig.json file
func (c Config) Version() string {
	return c.config.Version
}
