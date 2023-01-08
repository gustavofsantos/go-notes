package config

import (
	"os"
	"path"
)

type Config struct {
	path string
}

const DEFAULT_NOTES_DIR = "notes/"

func GetConfig() *Config {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultPath := path.Join(homedir, DEFAULT_NOTES_DIR)

	return &Config{path: defaultPath}
}

func GetPath(config *Config) string {
	return config.path
}
