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
        Port     string
        LogLevel string
    }

    PG struct {
        URL     string
        Host    string
        Port    string
        User    string
        Name    string
        Passwd  string
        MaxConn int
        Sslmode string
    }

    Log struct {
        Level string
    }
)

func NewConfig(config string) (*Config, error) {
    cfg := &Config{}

    if config == "default" {
        cfg.PG.Host = "db"
        cfg.PG.Port = "5432"
        cfg.PG.User = "postgres"
        cfg.PG.Name = "db"
        cfg.PG.Passwd = "password"
        cfg.PG.Sslmode = "disable"
        cfg.PG.URL = getDbUrl(cfg)
        cfg.PG.MaxConn = 10
        cfg.Log.Level = "ERROR"
        cfg.HTTP.Port = "8080"
        cfg.HTTP.LogLevel = "ERROR"
        return cfg, nil
    }

    err := godotenv.Load(config)
    if err != nil {
        return nil, fmt.Errorf("config error: %w", err)
    }

    cfg.PG.Host = os.Getenv("POSTGRES_HOST")
    cfg.PG.Port = os.Getenv("POSTGRES_PORT")
    cfg.PG.User = os.Getenv("POSTGRES_USER")
    cfg.PG.Name = os.Getenv("POSTGRES_NAME")
    cfg.PG.Passwd = os.Getenv("POSTGRES_PASSWORD")
    cfg.PG.Sslmode = os.Getenv("SSLMODE")
    cfg.PG.URL = getDbUrl(cfg)
    cfg.PG.MaxConn, _ = strconv.Atoi(os.Getenv("DB_MAXCONNS"))
    cfg.Log.Level = os.Getenv("LOG_LEVEL")
    cfg.HTTP.Port = os.Getenv("PORT")
    cfg.HTTP.LogLevel = cfg.Log.Level

    return cfg, nil
}

func getDbUrl(cfg *Config) string {
    port, err := strconv.Atoi(cfg.PG.Port)
    if err != nil {
        port = 5432
    }
    url := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=%s",
        cfg.PG.Host, port, cfg.PG.User, cfg.PG.Passwd, cfg.PG.Name, cfg.PG.Sslmode)
    return url
}
