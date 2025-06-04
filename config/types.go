package config

type Config struct {
	Redis struct {
		URL      string `yaml:"url" optional:"true"`
		Password string `yaml:"password" optional:"true"`
	} `yaml:"redis"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
