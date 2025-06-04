package config

type Config struct {
	Redis struct {
		URL      string `optional:"true" yaml:"url"`
		Password string `optional:"true" yaml:"password"`
	} `yaml:"redis"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
