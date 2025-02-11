package config

import (
	"bytes"
	"encoding/json"
	"strings"
)

// UnmarshalJSON provides a custom unmarshaling implementation for Config.
func (c *Config) UnmarshalJSON(data []byte) error {
	result := bytes.NewReader(data)
	decoder := json.NewDecoder(result)
	err := decoder.Decode(&c.config)

	return err
}

// MarshalJSON provides a custom marshaling implementation for Config.
func (c Config) MarshalJSON() ([]byte, error) {
	result := strings.Builder{}
	encoder := json.NewEncoder(&result)
	err := encoder.Encode(c.config)

	return []byte(result.String()), err
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
