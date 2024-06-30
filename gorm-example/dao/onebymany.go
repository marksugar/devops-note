package dao

import (
	"fmt"
	"log"
	"log/slog"
	"orm/db"

	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Name   string
	MoneID int
}
type Mone struct {
	gorm.Model
	Name string
	Cats []Cat
}

func (d Dao) OnebyManyCreate() {
	d.DB().AutoMigrate(&Cat{}, &Mone{})
	c := Cat{
		Name: "布兰奇",
	}
	c1 := Cat{
		Name: "踏雪",
	}
	m := Mone{
		Name: "cc",
		Cats: []Cat{c, c1},
	}
	d.DB().Create(&m)
}

// https://gorm.io/zh_CN/docs/has_many.html
func (d Dao) OnebyManySelect() {

	// 查询所有
	var m []Mone
	err := d.DB().Preload("Cats").Model(m).Where("name = ?", "cc").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m)

	// 查询单条cats的
	err = d.DB().Preload("Cats", "name = ?", "踏雪").Model(m).Where("name = ?", "cc").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m)

	// https://gorm.io/zh_CN/docs/preload.html#%E8%87%AA%E5%AE%9A%E4%B9%89%E9%A2%84%E5%8A%A0%E8%BD%BD-SQL
	err = d.DB().Preload("Cats", func(db *gorm.DB) *gorm.DB {
		return db.Order("test_cats.name DESC")
	}).Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println("3:", m)

	err = d.DB().Preload("Cats", func(db *gorm.DB) *gorm.DB {
		return db.Order("test_cats.name DESC")
	}).Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println("3:", m)
}

type Ect2 struct {
	gorm.Model
	Name   string
	Type   string
	Cat2ID uint
}
type Cat2 struct {
	gorm.Model
	Name    string
	Mone2ID int
	Ect2    []Ect2
}
type Mone2 struct {
	gorm.Model
	Name string
	Cats []Cat2
}

func (d Dao) OnebyManyCreate2() {

	d.DB().AutoMigrate(&Ect2{}, &Cat2{}, &Mone2{})
	e1 := Ect2{
		Name: "渴望",
		Type: "鸡肉",
	}
	e2 := Ect2{
		Name: "爱肯拿",
		Type: "鱼肉",
	}
	c1 := Cat2{
		Name: "布兰奇",
		Ect2: []Ect2{e1, e2},
	}
	c2 := Cat2{
		Name: "踏雪",
	}
	m := Mone2{
		Name: "cc",
		Cats: []Cat2{c1, c2},
	}
	d.DB().Create(&m)

	// mone := Mone2{Name: "Mone Example"}
	// cat := Cat2{Name: "Cat Example", Ect2: []Ect2{{Name: "Ect Example", Type: "Type Example"}, {Name: "Ect2 Example", Type: "Type2 Example"}}}
	// mone.Cats = append(mone.Cats, cat)
	// d.DB().Create(&mone)
}
func (d Dao) OnebyManySelect2() {

	// 查询所有
	var m []Mone2
	err := d.DB().Preload("Cats.Ect2").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m, "\n==============\n")

	err = d.DB().Preload("Cats.Ect2", "name = ?", "渴望").Preload("Cats").Model(&m).Where("name = ?", "cc").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m, "\n==============\n")

	err = d.DB().Preload("Cats.Ect2", "name = ?", "渴望").Preload("Cats").Model(&m).Where("name = ?", "cc").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m, "\n==============\n")

	err = d.DB().Preload("Cats.Ect2", "name = ?", "渴望").Preload("Cats", "name = ?", "布兰奇").Model(&m).Where("name = ?", "cc").Find(&m).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m, "\n==============\n")

}

type Ect22 struct {
	gorm.Model
	Name    string
	Type    string
	Cat22ID uint
}

type Cat22 struct {
	gorm.Model
	Name  string
	Ect22 []Ect22
}

type Mone22 struct {
	gorm.Model
	Name string
	Cats []Cat22 `gorm:"many2many:mone22_cats;"`
}

func (d Dao) OnebyManyCreate3() {

	d.DB().AutoMigrate(&Ect22{}, &Cat22{}, &Mone22{})

	e1 := Ect22{
		Name: "渴望",
		Type: "鸡肉",
	}
	e2 := Ect22{
		Name: "爱肯拿",
		Type: "鱼肉",
	}

	// 创建 Cat2 数据
	c1 := Cat22{
		Name:  "布兰奇",
		Ect22: []Ect22{e1, e2},
	}
	c2 := Cat22{
		Name:  "踏雪",
		Ect22: []Ect22{e1, e2},
	}

	// 创建 Mone2 数据
	m := Mone22{
		Name: "cc",
		Cats: []Cat22{c1, c2},
	}
	// 使用事务插入数据
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&m).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to insert data: %v", err)
	}

}
func (d Dao) OnebyManySelect3() {
	// 查询所有
	var m22 []Mone22
	err := d.DB().Preload("Cats.Ect22").Find(&m22).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(m22, "\n==============\n")

	err = d.DB().Preload("Cats.Ect22", "name = ?", "渴望").Preload("Cats").Model(&m22).Where("name = ?", "cc").Find(&m22).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Print(m22, "\n==============\n")

	err = d.DB().Preload("Cats.Ect22", "name = ?", "渴望").Preload("Cats").Model(&m22).Where("name = ?", "cc").Find(&m22).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Print(m22, "\n==============\n")

	err = d.DB().Preload("Cats.Ect22", "name = ?", "渴望").Preload("Cats", "name = ?", "布兰奇").Model(&m22).Where("name = ?", "cc").Find(&m22).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Print(m22, "\n==============\n")

}
func (d Dao) OnebyManySelect4() {
	var mm22 []Mone22
	// https://gorm.io/zh_CN/docs/preload.html#Joins-%E9%A2%84%E5%8A%A0%E8%BD%BD

	// err := d.DB().Preload("Cats", func(db *gorm.DB) *gorm.DB {
	// 	return db.Joins("Ect22").Where(Ect22{Name: "js"})
	// }).Where("name = ? ", "ww").Find(&mm22).Error
	err := d.DB().Preload("Cats.Ect22").Where("name = ? ", "ww").Find(&mm22).Error
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Print(mm22, "\n======Joins1========\n")
}
