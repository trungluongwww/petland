package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/config"
	"log"
)

func init() {
	config.Init()
}

func main() {
	env := config.GetEnv()

	e := echo.New()

	log.Panic(e.Start(fmt.Sprintf(":%s", env.AppPort)))
}
