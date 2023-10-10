package config

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host, Database, User, Password string
	Port                           int
}
