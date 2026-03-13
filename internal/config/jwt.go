package config

type JWTConfig struct {
	Secret          string
	ExpirationHours int
}
