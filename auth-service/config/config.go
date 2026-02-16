package config

import "os"

type Config struct {
	Port      string
	JWTSecret string
}

func Load() Config {
	return Config{
		Port:      getEnv("PORT", "8001"),
		JWTSecret: getEnv("JWT_SECRET", "supersecret"),
	}
}

func getEnv(key, fallback string) string {

	// it is func in os package that checks for the env varible
	//  in os with the key and returns the value as string
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
