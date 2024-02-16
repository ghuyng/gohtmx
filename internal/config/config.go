package config

type Config struct {
	Server ServerConfig `mapstructure:",squash"`
}

type ServerConfig struct {
	Addr                       string   `mapstructure:"SERVER_ADDR"`
	ReadTimeoutInSeconds       int      `mapstructure:"SERVER_READ_TIMEOUT_IN_SECONDS"`
	ReadHeaderTimeoutInSeconds int      `mapstructure:"SERVER_READ_HEADER_TIMEOUT_IN_SECONDS"`
	WriteTimeoutInSeconds      int      `mapstructure:"SERVER_WRITE_TIMEOUT_IN_SECONDS"`
	IdleTimeoutInSeconds       int      `mapstructure:"SERVER_IDLE_TIMEOUT_IN_SECONDS"`
	AllowOrigins               []string `mapstructure:"SERVER_ALLOW_ORIGINS"`
	AllowMethods               []string `mapstructure:"SERVER_ALLOW_METHODS"`
}
