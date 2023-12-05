package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trungluongwww/petland/config"
	v1 "github.com/trungluongwww/petland/pkg/router/v1"
)

const (
	versionOne = "/v1"
)

func Init(e *echo.Echo, env config.Env) {
	// default middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/api")

	// version 1
	gV1 := r.Group(versionOne)

	v1.Init(gV1)
}
