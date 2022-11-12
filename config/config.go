package config

type Config struct {
	ServerConfig `mapstructure:",squash"`
	MySQLConfig  `mapstructure:",squash"`
}

type ServerConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}
