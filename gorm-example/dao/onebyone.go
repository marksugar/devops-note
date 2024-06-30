package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name     string
	MasterID uint
	Master   Master
}
type Master struct {
	gorm.Model
	Name string
}

func (d Dao) Belongs_to() {
	// Belongs_to https://gorm.io/zh_CN/docs/belongs_to.html

	ms := Master{
		Name: "袅袅",
	}
	dg := Dog{
		Name:   "nike",
		Master: ms,
	}

	d.DB().AutoMigrate(&Dog{})
	d.DB().Create(&dg)
}

// has one
type Dog1 struct {
	gorm.Model
	Name      string
	Master1ID uint
}
type Master1 struct {
	gorm.Model
	Name string
	Dog1 Dog1
}

func (d Dao) Has_one() {
	// 1
	dg1 := Dog1{
		Name: "nike",
	}
	ms1 := Master1{
		Name: "袅袅",
		Dog1: dg1,
	}

	d.DB().AutoMigrate(&Master1{}, &Dog1{})
	d.DB().Create(&ms1)

	// 2
	// Preload 预加载
	// https://gorm.io/zh_CN/docs/preload.html
	var ms Master1
	d.DB().Preload("Dog1").First(&ms, 1)

	fmt.Println(ms)

	// 3
	// has one

	dg2 := Dog2{
		Name: "nike",
	}
	ms2 := Master2{
		Name: "袅袅",
		Dog2: dg2,
	}

	d.DB().AutoMigrate(&Master2{}, &Dog2{})
	// d.DB().Create(&ms2)
	// d.DB().Create(&dg2)
	d.DB().Model(&Dog2{}).Association("Master2").Append(&ms2)
}

type Dog2 struct {
	gorm.Model
	Name      string
	Master2ID uint
}
type Master2 struct {
	gorm.Model
	Name string
	Dog2 Dog2
}
