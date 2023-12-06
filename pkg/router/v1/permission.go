package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/pkg/handler"
)

func permission(g *echo.Group) {
	var (
		r = g.Group("/permission")
		h = handler.NewPermission()
	)

	r.POST("", h.Create)
}
