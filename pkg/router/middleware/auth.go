package routermiddleware

import (
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/config"
	"github.com/trungluongwww/petland/internal/utils/response"
	"log"
)

type Auth interface {
	RequiredAdmin() echo.MiddlewareFunc
	RequiredLogin() echo.MiddlewareFunc
}

type auth struct {
	appSecretKey   string
	adminSecretKey string
}

func NewAuth(cfg config.Env) Auth {
	if !cfg.IsDevelop() {
		if cfg.SecretJwtAdmin == "" || cfg.SecretJwtApp == "" {
			log.Panicf(
				fmt.Sprintf("Secret key not found admin='%s', app='%s'",
					cfg.SecretJwtAdmin,
					cfg.SecretJwtApp))
		}
	}
	return &auth{
		appSecretKey:   cfg.SecretJwtApp,
		adminSecretKey: cfg.SecretJwtAdmin,
	}
}

func (m *auth) RequiredAdmin() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return response.R401(c, nil, "")
			}

			return nil
		},
		ContextKey: "user",
		SigningKey: []byte(m.adminSecretKey),
	})
}

func (m *auth) RequiredLogin() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return response.R401(c, nil, "")
			}

			return nil
		},
		ContextKey: "user",
		SigningKey: []byte(m.appSecretKey),
	})
}
