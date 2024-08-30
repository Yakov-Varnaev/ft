package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	POSTGRES_HOST_ENV     = "POSTGRES_HOST"
	POSTGRES_PORT_ENV     = "POSTGRES_PORT"
	POSTGRES_DATABASE_ENV = "POSTGRES_DATABASE"
	POSTGRES_USER_ENV     = "POSTGRES_USER"
	POSTGRES_PASSWORD_ENV = "POSTGRES_PASSWORD"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func NewDatabase() (PostgresConfig, error) {
	cfg := PostgresConfig{}
	var err, envError error

	cfg.Host, err = readStrEnv(POSTGRES_HOST_ENV, "localhost")
	envError = errors.Join(envError, err)

	cfg.Port, err = readIntEnv(POSTGRES_PORT_ENV, "5432")
	envError = errors.Join(envError, err)

	cfg.Database, err = readStrEnv(POSTGRES_DATABASE_ENV, "ft")
	envError = errors.Join(envError, err)

	cfg.User, err = readStrEnv(POSTGRES_USER_ENV, "ft_user")
	envError = errors.Join(envError, err)

	cfg.Password, err = readStrEnv(POSTGRES_PASSWORD_ENV, "password")
	envError = errors.Join(envError, err)

	return cfg, envError
}

func (c *PostgresConfig) AsConnString() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Database, c.User, c.Password,
	)
}

type Config struct {
	DB PostgresConfig
}

func New() (Config, error) {
	cfg := Config{}
	var err, envError error

	cfg.DB, err = NewDatabase()
	envError = errors.Join(envError, err)

	return cfg, envError
}

func readStrEnv(key string, def ...string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		if len(def) > 0 {
			return def[0], nil
		}
	}
	return "", fmt.Errorf("env var not set: %s", key)
}

func readIntEnv(key string, def ...string) (int, error) {
	value, err := readStrEnv(key, def...)
	if err != nil {
		return 0, err
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
