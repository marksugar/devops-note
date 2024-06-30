package dao

import (
	"errors"
	"fmt"
	"log/slog"
	"orm/models"

	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/update.html
// 单字段，多字段
func (d Dao) Update() {

	// 单字段
	// update 只更新选择字段
	var u models.UserInfo
	res1 := d.DB().Model(&u).Where(models.UserInfo{Name: "justin"}).Update("name", "justins")
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res1 查询失败")
	}
	fmt.Print(u, "\n")
	fmt.Println("====")

	// save 无论如何都更新，所有内容，包括0
	// 如果没有主键就会创建新的
	// 根据主键更新、把符合主键的信息的结果中的内容都更新
	var us []models.UserInfo
	res3 := d.DB().Model(&us).Where(models.UserInfo{Name: "danny"}).Find(&us)
	// 将所有danny的age改成100
	for k := range us {
		us[k].Age = 1001
	}
	res3.Save(us)
	if errors.Is(res3.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res3 查询失败")
	}
	fmt.Print(us, "\n")
	fmt.Println("====")

	// updates 更新所有字段，一种map，一种结构体，结构体零值不更新

	// map可以更新0值
	// 更新单条
	var da3 models.UserInfo
	// 更新多条
	// var da3 []models.UserInfo
	res4 := d.DB().Model(&da3).Where(models.UserInfo{Name: "justin", Age: 23}).Updates(map[string]interface{}{"Name": "justins", "Age": 0})
	if errors.Is(res4.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res2 查询失败")
	}
	fmt.Print(u, "\n")
	fmt.Println("====")

	// struct
	// 结构体不能更新0值
	// 多条
	// type data2 struct {
	// 	Name string
	// 	age  int
	// }
	var da2 []models.UserInfo
	res2 := d.DB().Model(&da2).Where(models.UserInfo{Name: "justins"}).Updates(models.UserInfo{Name: "justin", Age: 23})
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res2 查询失败")
	}
	fmt.Print(u, "\n")
	fmt.Println("====")

}
