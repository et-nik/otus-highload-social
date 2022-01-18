package config

import (
	"errors"
	"os"
	"strconv"
)

var errEmptyPort = errors.New("empty port")
var errEmptyDatabase = errors.New("empty database")

type Config struct {
	Port     uint16
	Database string
}

func LoadConfig() (*Config, error) {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		return nil, errEmptyPort
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return nil, err
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return nil, errEmptyDatabase
	}

	return &Config{
		Port:     uint16(port),
		Database: database,
	}, nil
}
