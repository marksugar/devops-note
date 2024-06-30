package dao

import (
	"orm/models"

	"gorm.io/gorm"
)

// omit
// https://gorm.io/zh_CN/docs/create.html#%E7%94%A8%E6%8C%87%E5%AE%9A%E7%9A%84%E5%AD%97%E6%AE%B5%E5%88%9B%E5%BB%BA%E8%AE%B0%E5%BD%95

func (d Dao) Insert(uinfo *[]models.UserInfo) *gorm.DB {
	// res := d.db.Create(uinfo)
	// fmt.Println(res.Error, res.RowsAffected)

	return d.db.Create(uinfo)
}

// 只创建name
func (d Dao) InsertNameOne(uinfo *models.UserInfo) *gorm.DB {
	// res := d.db.Create(uinfo)
	// fmt.Println(res.Error, res.RowsAffected)

	return d.db.Select("Name").Create(uinfo)
}

// 查询
func (d Dao) Select(uinfo *[]models.UserInfo) *gorm.DB {
	// res := d.db.Create(uinfo)
	// fmt.Println(res.Error, res.RowsAffected)

	return d.db.Create(uinfo)
}
