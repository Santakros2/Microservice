package config

type Config struct {
	AuthSecret string
	ProfileURL string
	AuthURL    string
}

func Load() *Config {
	return &Config{
		AuthSecret: "supersecret", // later from env
		ProfileURL: "http://profile:8002",
		AuthURL:    "http://auth:8001",
	}
}
