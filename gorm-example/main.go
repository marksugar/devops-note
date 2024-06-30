package main

import (
	"fmt"
	"log/slog"
	"orm/dao"
	"orm/db"
	"orm/models"
)

func Insert(d *dao.Dao) {
	user := &[]models.UserInfo{
		{Name: "edinw", Age: 23},
		{Name: "justin", Age: 23},
		{Name: "boom", Age: 23},
		{Name: "nike", Age: 23},
		{Name: "danny", Age: 23}}
	// 1.插入结构体
	res := d.Insert(user)
	if res.Error != nil {
		slog.Error("Insert插入失败")
	}
	fmt.Printf("Insert生效了%d行\n", res.RowsAffected)

	// 2.单行插入
	res2 := d.InsertNameOne(&models.UserInfo{Name: "sean"})
	if res2.Error != nil {
		slog.Error("InsertNameOne插入失败", res2.Error)
	}
	fmt.Printf("Insert生效了%d行\n", res2.RowsAffected)
}

// create database orm_test charset utf8mb4
func main() {
	d, err := dao.NewDao()
	if err != nil {
		slog.Info(fmt.Sprintf("err:%s", err))
	}
	// // 数据插入
	// Insert(d)
	// d.Find()
	// d.First()
	// d.FindOne()
	// d.Cutdata()
	// d.Update()
	// d.Delete()
	// d.Scan()
	// d.Belongs_to()
	// d.Has_one()
	// d.OnebyManyCreate()
	// d.OnebyManySelect()
	// d.OnebyManyCreate2()
	// d.OnebyManySelect2()

	// // 一对多的创建和查询
	// d.OnebyManyCreate3()
	// d.OnebyManySelect3()
	// d.OnebyManySelect4()
	// d.ManyByMany()

	// // 多对多的创建和查询和条件查询
	// d.Many2Create()
	// d.Many2SelectOne()
	// d.Many2SelectOneByOne()
	// // 多对多的id关联
	// d.Many2Link()

	// 多态标签
	// d.Polymorphism()

	// // 指定外键
	// d.AssociationTags()
	// 多对多指定外键
	// d.AssociationTagsMany2()
	// d.AssociationTagsMany2Select()

	// 事务
	// d.Transactions()
	// d.TransactionsBegin()
	// d.TransactionsBeginTo()

	// 自定义数据类型
	d.DataTypeString()
	db.Close(db.DB)
}
