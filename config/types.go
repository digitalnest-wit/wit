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
