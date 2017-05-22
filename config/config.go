package config

type ApplicationConfig struct {
	Port  int
	Mysql MysqlConfig
}

type MysqlConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}