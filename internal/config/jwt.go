package config

import "os"

func GetJWTSecret() []byte {
	secret_key := os.Getenv("JWT_SECRET")

	if secret_key == "" {
		panic("JWT Secret is not set")
	}

	return []byte(secret_key)
}
