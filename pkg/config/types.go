package config

type Config struct {
	Listener struct {
		Address string `yaml:"address"`
	} `yaml:"listener"`
	Logger struct {
		Level string `yaml:"level"`
	} `yaml:"logger"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
