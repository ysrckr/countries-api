package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ysrckr/countries-api/internals/server"
)

func main() {

	server := server.New()

	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
