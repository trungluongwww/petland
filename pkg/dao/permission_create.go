package dao

import (
	"github.com/trungluongwww/petland/internal/db"
	"github.com/trungluongwww/petland/internal/model/mysql/mysqlpetland"
	"gorm.io/gorm"
)

func (permission) InsertOne(gdb *gorm.DB, e *mysqlpetland.Permission) error {
	if gdb == nil {
		gdb = db.GetDB()
	}

	err := gdb.Save(e).Error
	if err != nil {
	}

	return err
}
