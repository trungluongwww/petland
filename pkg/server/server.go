package server

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/config"
	"github.com/trungluongwww/petland/internal/db"
	"github.com/trungluongwww/petland/internal/logger"
	"github.com/trungluongwww/petland/internal/utils/response"
	"github.com/trungluongwww/petland/pkg/router"
)

func Boostrap(e *echo.Echo, env config.Env) {
	// database
	db.Connect(env)

	logger.Init()

	response.Init()

	router.Init(e, env)
}
