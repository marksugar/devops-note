package dao

import (
	"orm/db"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

// NewDao 创建连接
func NewDao() (d *Dao, err error) {
	db := db.Initdb()
	d = &Dao{
		db: db,
	}
	return
}

func (d *Dao) DB() *gorm.DB {
	return d.db
}

func (d *Dao) Close() error {
	return db.Close(d.db)
}
