package config

type Config struct {
	Logger struct {
		Level string `yaml:"level"`
	} `yaml:"logger"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
