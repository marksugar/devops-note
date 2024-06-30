package dao

import (
	"orm/db"
	"orm/models"
)

// https://gorm.io/zh_CN/docs/delete.html
// 如果存在deleted_at字段，那么就会逻辑删除，也就是软删除
func (d *Dao) Delete() {
	// 逻辑删除
	var Last models.UserInfo
	db.DB.Model(&models.UserInfo{}).Where("name = ?", "boom").Delete(&Last)

	// Unscoped 可以跳过字段约束，直接删除。或者查询
	// https://gorm.io/zh_CN/docs/delete.html#%E6%9F%A5%E6%89%BE%E8%A2%AB%E8%BD%AF%E5%88%A0%E9%99%A4%E7%9A%84%E8%AE%B0%E5%BD%95
	var Last2 models.UserInfo
	db.DB.Model(&models.UserInfo{}).Unscoped().Where("name = ?", "boom").Delete(&Last2)
}
