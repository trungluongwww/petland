package requestmodel

import (
	"github.com/go-playground/validator/v10"
	"github.com/trungluongwww/petland/internal/model/mysql/mysqlpetland"
)

type PermissionCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (r PermissionCreate) Validate() error {
	return validator.New().Struct(&r)
}

func (r PermissionCreate) ConvertToEntity() *mysqlpetland.Permission {
	return &mysqlpetland.Permission{
		Name:        r.Name,
		Description: r.Description,
	}
}
