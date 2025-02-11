package config

// Config contains all the configuration data defined in witconfig.json
type Config struct {
	config witConfig
}

type Installer interface {
	Install() error
}

type witConfig struct {
	Version    string     `json:"version"`
	BrewConfig brewConfig `json:"brew"`
	CodeConfig codeConfig `json:"code"`
}

type brewConfig struct {
	Formulae []string `json:"formulae"`
	Casks    []string `json:"casks"`
}

type codeConfig struct {
	Extensions []string `json:"extensions"`
}
