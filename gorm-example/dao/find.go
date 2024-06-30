package dao

import (
	"errors"
	"fmt"
	"log/slog"
	"orm/models"

	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/query.html#%E6%A3%80%E7%B4%A2%E5%8D%95%E4%B8%AA%E5%AF%B9%E8%B1%A1
// 必须是一个模型才能查询
// First and Last 方法会按主键排序找到第一条记录和最后一条记录。first按照主键排序 take不会
func (d *Dao) First() {
	// var result map[string]interface{}
	// d.DB().Model(&models.UserInfo{}).First(&result)

	var First models.UserInfo
	d.DB().Model(&models.UserInfo{}).First(&First)

	var Last models.UserInfo
	d.DB().Model(&models.UserInfo{}).Last(&Last)

	var take models.UserInfo
	d.DB().Model(&models.UserInfo{}).Last(&take)
	fmt.Println(First, Last, take)

	// 主键检索，可以是主键id的检索
	var First2 models.UserInfo
	res := d.DB().Model(&models.UserInfo{}).First(&First2, 14)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		slog.Warn("查询失败")
	}
	fmt.Println(First2)

	var First3 models.UserInfo
	res3 := d.DB().Model(&models.UserInfo{}).First(&First3, "id =?", 14)
	if errors.Is(res3.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res3 查询失败")
	}
	fmt.Println(First3)

	var First4 models.UserInfo
	res4 := d.DB().Model(&models.UserInfo{}).First(&First4, map[string]interface{}{
		"name": "nike",
		"age":  23,
	})
	if errors.Is(res4.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res4 查询失败")
	}
	fmt.Println(First4)

	var First5 models.UserInfo
	res5 := d.DB().Model(&models.UserInfo{}).First(&First5, models.UserInfo{Name: "nike", Age: 23})
	if errors.Is(res5.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res5 查询失败")
	}
	fmt.Println(First5)
	// where查询的三种方式
	var whh models.UserInfo
	// 1,
	// whrere := d.DB().Where("name = ?", "nikes").First(&whh)
	// 2,
	// whrere := d.DB().Where(models.UserInfo{Name: "nike"}).First(&whh)
	// 3,
	whrere := d.DB().Where(map[string]interface{}{
		"name": "nikes",
	}).First(&whh)
	if errors.Is(whrere.Error, gorm.ErrRecordNotFound) {
		slog.Warn("whh查询失败")
	}
	fmt.Println(whh)

	// and or 查询
	// 满足其中一个
	and := d.DB().Where("name = ? And age = ?", "nikes", 18).Or("name = ? ", "nike").First(&whh)
	if errors.Is(and.Error, gorm.ErrRecordNotFound) {
		slog.Warn("and查询失败")
		return
	}
	fmt.Println(whh)

}

// 多条查询返回，和条件查询，模糊查询
func (d *Dao) Find() {
	var finds []models.UserInfo
	res1 := d.DB().Model(&models.UserInfo{}).Where(models.UserInfo{Name: "edinw", Age: 23}).Find(&finds)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res1 查询失败")
	}
	fmt.Print(finds, "\n")
	fmt.Println("====")

	res2 := d.DB().Model(&models.UserInfo{}).Where("name <> ?", "edinw").Find(&finds)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res2 查询失败")
	}
	fmt.Println(finds)
	fmt.Println("====")

	res3 := d.DB().Model(&models.UserInfo{}).Where("name LIKE ?", "%ea%").Find(&finds)
	if errors.Is(res3.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res3 查询失败")
	}
	fmt.Println(finds)
}

// 查询返回一个字段Select
// 不返回某一个字段Omit
func (d *Dao) FindOne() {
	var finds []models.UserInfo
	res1 := d.DB().Model(&models.UserInfo{}).Select("name").Where(models.UserInfo{Name: "edinw", Age: 23}).Find(&finds)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res1 查询失败")
	}
	fmt.Print(finds, "\n")
	fmt.Println("====")

	res2 := d.DB().Model(&models.UserInfo{}).Omit("name", "age").Where(models.UserInfo{Name: "edinw", Age: 23}).Find(&finds)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res2 查询失败")
	}
	fmt.Print(finds, "\n")
	fmt.Println("====")
}

// 到处固定的字段
func (d *Dao) Cutdata() {
	// 单条
	var data = struct {
		Name string
		age  int
	}{}
	res1 := d.DB().Model(&models.UserInfo{}).Where(models.UserInfo{Name: "edinw", Age: 23}).Find(&data)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res1 查询失败")
	}
	fmt.Print(data, "\n")
	fmt.Println("====")

	// 多条
	type data2 struct {
		Name string
		age  int
	}
	var da2 []data2
	res2 := d.DB().Model(&models.UserInfo{}).Where(models.UserInfo{Name: "edinw", Age: 23}).Find(&da2)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		slog.Warn(" res1 查询失败")
	}
	fmt.Print(da2, "\n")
	fmt.Println("====")
}
