package config

type Config struct {
	path string
}

func GetConfig() Config {
	return Config{path: "/home/gustavo/notes/"}
}
