package conf

import (
	"log"

	"github.com/joho/godotenv"
)

var mode = "development"

var Config = New(mode)

type Enviroment map[string]string

type Configuration struct {
	Envs    Enviroment
	mode    string
	envFile string
}

func New(mode string) *Configuration {
	config := &Configuration{}
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
