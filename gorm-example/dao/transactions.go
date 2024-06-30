package dao

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/transactions.html

type TS struct {
	ID   int
	Name string
}

func (d Dao) Transactions() {
	// 如果isok是false则不会插入数据
	// 如果isok是true，插入test1,2,3,并查询
	var isok bool // 默认false
	isok = true
	d.DB().AutoMigrate(&TS{})
	d.DB().Transaction(func(txS *gorm.DB) error {
		txS.Create(&TS{Name: "test1"})
		txS.Create(&TS{Name: "test2"})
		txS.Create(&TS{Name: "test3"})

		txS.Transaction(func(tx *gorm.DB) error {
			// tx.Create(&TS{Name: "test1bbb"})
			var tt []TS
			tx.Model(&tt).Find(&tt)
			fmt.Println(tt)
			if !isok {
				return errors.New("")
			}
			return nil

		})
		// 如果isok等于true，返回nil进行创建
		if isok {
			return nil
		}
		return errors.New("")

	})
}

// 手动提交
// commit之前的提交
// 如果isok等于true，就回滚commit之前的提交
func (d Dao) TransactionsBegin() {
	var isok bool // 默认false
	isok = true
	tx := d.DB().Begin()
	tx.Create(&TS{Name: "mark1"})
	tx.Create(&TS{Name: "mark2"})
	tx.Create(&TS{Name: "mark3"})
	if isok {
		tx.Rollback()
	}
	tx.Commit()
}

// 如果isok等于true,就跳转到save的bool点之前。之后的不提交
func (d Dao) TransactionsBeginTo() {
	var isok bool // 默认false
	isok = true
	tx := d.DB().Begin()
	tx.Create(&TS{Name: "mark1"})
	tx.Create(&TS{Name: "mark2"})
	tx.Create(&TS{Name: "mark3"})
	tx.SavePoint("bool")
	tx.Create(&TS{Name: "bool"})
	if isok {
		tx.RollbackTo("bool")
	}
	tx.Commit()
}
