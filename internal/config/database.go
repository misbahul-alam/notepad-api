package config

import "fmt"

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func (d *DatabaseConfig) DSN() string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", d.Username, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
	return url
}
