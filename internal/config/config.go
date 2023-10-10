package config

import "time"

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host, Database, User, Password string
	Port                           int
	MaxConns                       int32
	MaxConnsLifetime               time.Duration
}
