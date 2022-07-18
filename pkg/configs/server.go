package configs

import "fmt"

const (
	serverHost string = "localhost"
	serverPort string = "8080"
)

type serverConfigs struct {
	Addr string
}

var ServerCfg = serverConfigs{
	Addr: fmt.Sprintf("%s:%s", serverHost, serverPort),
}
