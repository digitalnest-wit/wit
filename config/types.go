package config

type Installer interface {
	Install() error
}
