package conf

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/server"
)

type Enviroment map[string]string

type Configuration struct {
	Envs    Enviroment
	Server  server.Server
	DB      db.Service
	mode    string
	envFile string
}

func New(mode string) *Configuration {
	config := &Configuration{}

	config.DB = db.New()
	config.Server = server.New()
	config.Envs = map[string]string{}
	config.mode = mode
	config.envFile = ".env"

	switch config.mode {
	case "development", "production", "ci", "test":
		config.setEnvFile()
	default:
		break
	}

	if err := config.setEnvs(); err != nil {
		log.Fatalf("cannot load %s file, got an error of %w", config.envFile, err)
	}

	return config
}

func (c *Configuration) setEnvFile() {
	c.envFile += "." + c.mode
}

func (c *Configuration) setEnvs() error {
	path := c.envFile
	var err error

	c.Envs, err = godotenv.Read(path)
	if err != nil {
		return err
	}

	return nil
}
