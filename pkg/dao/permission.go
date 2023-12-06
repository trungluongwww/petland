package dao

import (
	"github.com/trungluongwww/petland/internal/model/mysql/mysqlpetland"
	"gorm.io/gorm"
)

type Permission interface {
	InsertOne(gdb *gorm.DB, e *mysqlpetland.Permission) error
	FirstByName(gdb *gorm.DB, name string) (*mysqlpetland.Permission, error)
}

type permission struct {
}

func NewPermission() Permission {
	return &permission{}
}
