package dao

import (
	"fmt"
	"orm/db"
	"orm/models"
)

// https://gorm.io/zh_CN/docs/sql_builder.html
func (d Dao) Scan() {
	var u []models.UserInfo
	db.DB.Raw("select uuid,name,age from test_user_infos where name = ?", "sean").Scan(&u)
	fmt.Println(u)
}
