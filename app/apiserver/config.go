package apiserver

// Config ...
type Config struct {
	BindAddr       string `toml:"bind_addr"`
	DatabaseDriver string `toml:"database_driver"`
	DatabaseURL    string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
