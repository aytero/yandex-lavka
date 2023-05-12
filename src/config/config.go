package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	HTTP
	PG
	Log
}

type (
	HTTP struct {
		Port string
	}

	PG struct {
		URL     string
		Host    string
		Port    string
		User    string
		Name    string
		Passwd  string
		MaxConn int
		//PoolMax int
		// sslmode
		// connect timeout
	}

	Log struct {
		Level string
	}
)

func NewConfig(config string) (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load(config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	cfg.PG.Host = os.Getenv("POSTGRES_HOST")
	cfg.PG.Port = os.Getenv("POSTGRES_PORT")
	cfg.PG.User = os.Getenv("POSTGRES_USER")
	cfg.PG.Name = os.Getenv("POSTGRES_NAME")
	cfg.PG.Passwd = os.Getenv("POSTGRES_PASSWORD")
	cfg.PG.URL = getDbUrl(cfg)
	cfg.PG.MaxConn, _ = strconv.Atoi(os.Getenv("DB_MAXCONNS"))

	cfg.Log.Level = os.Getenv("LOG_LEVEL")

	return cfg, nil
}

func getDbUrl(cfg *Config) string {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		cfg.PG.Host, cfg.PG.Port, cfg.PG.User, cfg.PG.Passwd, cfg.PG.Name)

	return url
}
