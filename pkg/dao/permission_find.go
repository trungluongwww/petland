package dao

import (
	"github.com/trungluongwww/petland/internal/db"
	"github.com/trungluongwww/petland/internal/logger"
	"github.com/trungluongwww/petland/internal/model/mysql/mysqlpetland"
	"gorm.io/gorm"
)

func (permission) FirstByName(gdb *gorm.DB, name string) (*mysqlpetland.Permission, error) {
	var (
		found *mysqlpetland.Permission
	)

	if gdb == nil {
		gdb = db.GetDB()
	}

	err := gdb.First(&found, &mysqlpetland.Permission{Name: name}).Error
	if err != nil {
		logger.Error("dao.perrmission.FindOne", err, name)
		return nil, err
	}

	return found, nil
}
