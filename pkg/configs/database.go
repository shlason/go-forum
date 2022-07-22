package configs

import "fmt"

const (
	dialect  string = "mysql"
	username string = "root"
	password string = "password"
	protocol string = "tcp"
	dbHost   string = "localhost"
	dbName   string = "forum"
	options  string = "charset=utf8&parseTime=True&loc=Local"
)

type databaseConfigs struct {
	Dsn     string
	Dialect string
}

var DbCfg = databaseConfigs{
	Dsn:     fmt.Sprintf("%s:%s@%s(%s)/%s?%s", username, password, protocol, dbHost, dbName, options),
	Dialect: dialect,
}
