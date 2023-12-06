package service

import (
	"context"
	requestmodel "github.com/trungluongwww/petland/pkg/model/request"
)

type Permission interface {
	Create(ctx context.Context, p requestmodel.PermissionCreate) error
}

type permission struct{}

func NewPermission() Permission {
	return &permission{}
}
