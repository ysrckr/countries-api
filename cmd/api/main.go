package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ysrckr/countries-api/internals/conf"
)

func main() {

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
