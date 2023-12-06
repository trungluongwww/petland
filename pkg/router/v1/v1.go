package v1

import "github.com/labstack/echo/v4"

func Init(g *echo.Group) {
	r := g.Group("")

	permission(r)
}
