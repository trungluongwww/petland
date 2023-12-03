package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/config"
	"github.com/trungluongwww/petland/internal/db"
	"github.com/trungluongwww/petland/internal/model/mysql/mysqlpetland"
	"gorm.io/gorm"
	"log"
	"net/http"
)

//go:generate gentool -onlyModel -outPath ./internal/model/mysql/reviewer -modelPkgName msqlreviewer

func init() {
	config.Init()
}

func main() {
	env := config.GetEnv()

	dbc := db.Connect(env)

	e := echo.New()

	Router(e, dbc)

	log.Panic(e.Start(fmt.Sprintf(":%s", env.AppPort)))
}

func Router(e *echo.Echo, db *gorm.DB) {
	g := e.Group("")

	g.GET("/hello-world", func(c echo.Context) error {
		err := db.Save(&mysqlpetland.Permission{
			Name: "order_admin",
		}).Error

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"Hello": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"Hello": "world"})
	})
}
