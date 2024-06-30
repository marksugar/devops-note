package dao

import (
	"fmt"
	"orm/db"

	"gorm.io/gorm"
)

type Ect struct {
	gorm.Model
	Name  string
	Money uint
	// Cats  []Cat `gorm:"many2many:cat_ect;"`
}
type Ccat struct {
	gorm.Model
	Name string
	Age  int
	Ects []Ect `gorm:"many2many:cat_ect;"`
}
type People struct {
	gorm.Model
	Name string
	Cats []Ccat `gorm:"many2many:people_cat;"`
}

func (d Dao) Many2Create() {
	db.DB.AutoMigrate(&People{}, &Ccat{})
	e1 := Ect{
		Name:  "爱肯拿",
		Money: 487,
	}
	e2 := Ect{
		Name:  "渴望",
		Money: 390,
	}
	c1 := Ccat{
		Name: "踏雪",
		Age:  2,
		Ects: []Ect{e1, e2},
	}
	c2 := Ccat{
		Name: "布兰奇",
		Age:  2,
		Ects: []Ect{e1, e2},
	}
	p := People{
		Name: "sean",
		Cats: []Ccat{c1, c2},
	}
	db.DB.Create(&p)
}
func (d Dao) Many2SelectAll() {
	var pp []People
	db.DB.Preload("Cats.Ects").Model(&People{}).Find(&pp)
	fmt.Println(pp)
}

// 条件查询
func (d Dao) Many2SelectOne() {
	var pp []People
	db.DB.Preload("Cats.Ects", "name = ? ", "爱肯拿").Preload("Cats", "name = ?", "布兰奇").Model(&People{}).Where("name = ?", "sean").Find(&pp)
	fmt.Println(pp)

	db.DB.Preload("Cats.Ects", "money < ? ", 400).Preload("Cats", "name = ?", "布兰奇").Model(&People{}).Where("name = ?", "sean").Find(&pp)
	fmt.Println(pp)
}

// 条件查询只查询

func (d Dao) Many2SelectOneByOne() {
	// 根据people的id 只查询Ccat的数据
	c1 := People{
		Model: gorm.Model{
			ID: 1,
		},
	}
	var pp []Ccat
	db.DB.Model(&c1).Association("Cats").Find(&pp)
	fmt.Println(pp)

	fmt.Println("=============")
	// 根据people的id 查询Ccat的数据,并且预加载Ects中所关联的数据
	db.DB.Model(&c1).Preload("Ects").Association("Cats").Find(&pp)
	fmt.Println(pp)

	fmt.Println("=============")
	// 根据people的id 查询Ccat的数据,并且预加载Ects中所关联的数据，根据条件判断
	db.DB.Model(&c1).Preload("Ects", "name=?", "渴望").Association("Cats").Find(&pp)
	fmt.Println(pp)

	fmt.Println("=============")
	// 根据people的id 查询Ccat的数据,并且预加载Ects中所关联的数据，根据条件判断
	db.DB.Model(&c1).Preload("Ects", "name=?", "渴望").Association("Cats").Find(&pp)
	fmt.Println(pp)

	fmt.Println("=============")
	c2 := People{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "sean",
	}
	// 根据people的id 查询Ccat的数据,并且预加载Ects中所关联的数据，根据条件判断
	db.DB.Model(&c2).Preload("Ects", "name=?", "渴望").Association("Cats").Find(&pp)
	fmt.Println(pp)
}

func (d Dao) Many2Link() {
	// 向ccat表id2 关联一个People表id为2的
	p1 := People{
		Model: gorm.Model{
			ID: 2,
		},
	}
	c1 := Ccat{
		Model: gorm.Model{
			ID: 2,
		},
	}

	// 添加新的关联
	db.DB.Model(&c1).Association("People").Append(&p1)

	// // 替换所有关联
	// 	db.DB.Model(&c1).Association("People").Replace(&p1)

	// // 删除关联
	// db.DB.Model(&c1).Association("People").Delete(&p1)

	// var pp []People
	// db.DB.Preload("Cats.Ects").Preload("Cats").Model(&People{}).Where("name = ?", "mark").Find(&pp)
	// fmt.Println(pp)
}

// func main() {
// 	// create database orm_test charset utf8mb4
// 	dsn := "root:1234@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		slog.Error(fmt.Sprintf("%s", err))
// 	}
// 	gdb = db
// 	// err = db.AutoMigrate(&People{}, &Cat{})
// 	// slog.Error(fmt.Sprintf("%s", err))

// 	sselectOne()
// }
