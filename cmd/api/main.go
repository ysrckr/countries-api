package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ysrckr/countries-api/internals/conf"
)

func main() {

	config := conf.New()
	fApp := config.Server.GetApp()

	config.Server.RegisterRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := fApp.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
