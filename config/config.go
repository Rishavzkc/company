package config

type Config struct {
	Database    DatabaseConfigs
	ServiceHost string
}

type DatabaseConfigs struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func NewConfig() *Config {
	// Hard coding sensitive info for now. we can use env variables to replace this step
	return &Config{
		Database: DatabaseConfigs{
			Username: "root",
			Password: "Quest1234",
			Host:     "localhost",
			Port:     3306,
			Database: "company",
		},
		ServiceHost: "localhost:8080",
	}
}
