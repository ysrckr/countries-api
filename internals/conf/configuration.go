package conf

import (
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/server"
)

type Enviroment map[string]string

type Configuration struct {
	Envs   Enviroment
	Server server.Server
	DB     db.Service
}

func New() *Configuration {
	config := &Configuration{}

	config.DB = db.New()
	config.Server = server.New()
	config.Envs = map[string]string{}

	return config
}
