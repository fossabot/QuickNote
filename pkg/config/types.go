package config

type Config struct {
	Listener struct {
		Address string `yaml:"address"`
		Static  string `optional:"true" yaml:"static"`
	} `yaml:"listener"`
	TLS struct {
		Cert string `optional:"true" yaml:"cert"`
		Key  string `optional:"true" yaml:"key"`
	} `yaml:"tls"`
	Logger struct {
		Level string `optional:"true" yaml:"level"`
		Dir   string `yaml:"dir"`
	} `yaml:"logger"`
	Database struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"database"`
}
