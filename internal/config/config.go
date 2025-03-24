package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	NAME string
	PORT string
	Env  string
	DB   *DatabaseConfig
}

func NewConfig() *Config {
	cfg := &Config{}
	return cfg.load()
}

func (c *Config) load() *Config {
	c.NAME = c.GetEnv("APP_NAME", "chat_go")
	c.PORT = c.GetEnv("PORT", "8080")
	c.Env = c.GetEnv("ENV", "development")

	c.DB = &DatabaseConfig{
		Host:     c.GetEnv("DB_HOST", "127.0.0.1"),
		Port:     c.GetEnv("DB_PORT", "5432"),
		User:     c.GetEnv("DB_USER", "chat_go"),
		Password: c.GetEnv("DB_PASSWORD", "password"),
		Name:     c.GetEnv("DB_NAME", "chat_go"),
	}

	return c
}

func (c *Config) GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
