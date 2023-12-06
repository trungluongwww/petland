package service

import (
	"context"
	"errors"
	"github.com/trungluongwww/petland/internal/db"
	"github.com/trungluongwww/petland/internal/utils/response"
	"github.com/trungluongwww/petland/pkg/dao"
	requestmodel "github.com/trungluongwww/petland/pkg/model/request"
)

func (permission) Create(ctx context.Context, p requestmodel.PermissionCreate) error {
	var (
		gdb = db.GetDB()
		d   = dao.NewPermission()
	)

	exists, _ := d.FirstByName(gdb, p.Name)
	if exists != nil {
		return errors.New(response.ErrAlreadyExists)
	}

	return d.InsertOne(gdb, p.ConvertToEntity())
}
