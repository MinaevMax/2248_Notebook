package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() (Config, error) {
	var cfg Config
	var err error

	cfg.Server, err = LoadServerConfig()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func LoadServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	var err error

	cfg.Port, err = getEnvInt("SERVER_PORT")
	if err != nil {
		return ServerConfig{}, err
	}

	cfg.ReadTimeout, err = getEnvDuration("SERVER_READ_TIMEOUT")
	if err != nil {
		return ServerConfig{}, err
	}

	cfg.WriteTimeout, err = getEnvDuration("SERVER_WRITE_TIMEOUT")
	if err != nil {
		return ServerConfig{}, err
	}

	return cfg, nil
}

func getEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env '%s' is not set", key)
	}
	return v, nil
}

func getEnvInt(key string) (int, error) {
	v, err := getEnv(key)
	if err != nil {
		return 0, err
	}
	intVal, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("atoi failed: %w", err)
	}
	return intVal, nil
}

func getEnvDuration(key string) (time.Duration, error) {
	v, err := getEnv(key)
	if err != nil {
		return 0, err
	}
	durationVal, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("parse duration failed: %w", err)
	}
	return durationVal, nil
}
