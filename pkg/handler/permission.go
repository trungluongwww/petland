package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/petland/internal/utils/response"
	requestmodel "github.com/trungluongwww/petland/pkg/model/request"
	"github.com/trungluongwww/petland/pkg/service"
)

type Permission interface {
	Create(c echo.Context) error
}

type permission struct{}

func NewPermission() Permission {
	return permission{}
}

func (permission) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		p   requestmodel.PermissionCreate
		s   = service.NewPermission()
	)

	if err := c.Bind(&p); err != nil {
		return response.R400(c, nil, err.Error())
	}

	if err := p.Validate(); err != nil {
		return response.RouterResponse(c, err)
	}

	if err := s.Create(ctx, p); err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, nil)
}
