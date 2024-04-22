package config

type HTTPServer struct {
	Port int `env:"SMD_SERVE_PORT"`
}

type Config struct {
	HTTPServer HTTPServer `env:"SMD_SERVE_HTTP"`
}
