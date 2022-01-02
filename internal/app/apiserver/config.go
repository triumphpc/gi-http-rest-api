package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	Loglevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// NewConfig instance nre config for server
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		Loglevel:    "debug",
		DatabaseURL: "",
	}
}
