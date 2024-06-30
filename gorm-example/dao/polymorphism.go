package dao

// 多态 https://gorm.io/zh_CN/docs/polymorphism.html
// 两个字段与多张表
// 一对多或者多对多
// `gorm:"polymorphic:Owner;"` Owner的ID和type是多态里面的关联字段

// 其他的标签：
// polymorphicType: 指定类型
// polymorphicId: 指定ID
// polymorphicValue: 指定类型的值,如果配置则会替换默认的结构体名称
type SonA struct {
	ID      int
	Name    string
	MasterA []MasterA `gorm:"polymorphic:Owner;"`
}
type SonB struct {
	ID      int
	Name    string
	MasterA MasterA `gorm:"polymorphic:Owner;"`
}
type MasterA struct {
	ID        int
	Name      string
	OwnerType string
	OwnerID   int
}

func (d Dao) Polymorphism() {
	// 创建表SonA和SonB，Master、
	// Master存放的是Master和SonA和SonB的ID

	d.DB().AutoMigrate(&SonA{}, &SonB{}, &MasterA{})

	// 创建小明和小师兄关联
	// SONA中小明的ID会与Master中创建的小师兄的OwnerID关联，OwnerType是SONA的表名
	d.DB().Create(&SonA{Name: "小明", MasterA: []MasterA{
		{Name: "小师兄"},
		{Name: "大师兄"},
	}})
	// 创建小红和大师兄关联
	// SONA中小红的ID会与Master中创建的大师兄的OwnerID关联，OwnerType是SONA的表名
	d.DB().Create(&SonB{Name: "小红", MasterA: MasterA{
		Name: "大师兄",
	}})
}

// mysql> select * from son_as;
// +----+------+
// | id | name |
// +----+------+
// |  1 | 小明 |
// +----+------+
// 1 row in set (0.03 sec)

// mysql> select * from son_bs;
// +----+------+
// | id | name |
// +----+------+
// |  1 | 小红 |
// +----+------+
// 1 row in set (0.04 sec)

// mysql> select * from master_as;
// +----+--------+------------+----------+
// | id | name   | owner_type | owner_id |
// +----+--------+------------+----------+
// |  1 | 小师兄 | son_as     |        1 |
// |  2 | 大师兄 | son_as     |        1 |
// |  3 | 大师兄 | son_bs     |        1 |
// +----+--------+------------+----------+
// 3 rows in set (0.05 sec)
