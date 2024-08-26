package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ysrckr/countries-api/internals/conf"
)

var env string

func init() {

	flag.StringVar(&env, "env-file", ".env", "Set env-file path")
	flag.StringVar(&env, "e", ".env", "Set env-file path")
	flag.Parse()

}

func main() {
	log.Println(env)
	config := conf.New()

	config.Server.RegisterRoutes()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}
	err = config.Server.StartServer(port)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
