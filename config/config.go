package config

import "github.com/BurntSushi/toml"

type Config struct {
	KeyWord []string
}

var config = &Config{
	KeyWord: []string{"中国"},
}

func LoadConfig() *Config {
	filePath := ""

	if _, err := toml.Decode(filePath, &config); err != nil {

	}

	return config
}
