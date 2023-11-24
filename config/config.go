package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RESTPort    string `envconfig:"REST_PORT" default:"8080"`
	MySqlDbName string `envconfig:"MYSQL_DB_NAME" default:"galactic"`
	MysqlDbHost string `envconfig:"MYSQL_DB_HOST" default:"localhost"`
	MysqlDbUser string `envconfig:"MYSQL_DB_USER" default:"root"`
	MysqlDbPass string `envconfig:"MYSQL_DB_PASS" default:"root"`
	MysqlDbPort string `envconfig:"MYSQL_DB_PORT" default:"3306"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
