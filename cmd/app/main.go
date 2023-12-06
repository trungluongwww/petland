package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/config"
	"github.com/trungluongwww/petland/pkg/server"
	"log"
)

func init() {
	config.Init()
}

func main() {
	env := config.GetEnv()

	e := echo.New()

	server.Boostrap(e, env)

	log.Panic(e.Start(fmt.Sprintf(":%s", env.AppPort)))
}
