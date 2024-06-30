package dao

import "fmt"

// https://gorm.io/zh_CN/docs/associations.html#%E5%85%B3%E8%81%94%E6%A0%87%E7%AD%BE%EF%BC%88Association-Tags%EF%BC%89
// 指定外键
// foreignKey: 指定另外一个表的字段为外键关联
// references： 重写这个关联的外键名称为当前指定结构体的字段名称
// 以下示例中的意思是说，AAS中MasterYI是一个切片，可以是多个，他的外键字段是MasterYI的AASName，而外键的值是AAS的Name,如果不指定references:Name。值是AAS的id
type AAS struct {
	ID       int
	Name     string
	MasterYI []MasterYI `gorm:"foreignKey:AASName;references:Name"`
}
type MasterYI struct {
	ID      int
	Name    string
	AASName string
}

//  `gorm:"foreignKey:AASName;references:Name"`
// 将MasterYI表中的AASName的内容是AAS.Name
// mysql> select * from master_yis;
// +----+--------+----------+
// | id | name   | aas_name |
// +----+--------+----------+
// |  1 | 石斑鱼 | 海洋     |
// |  2 | 海蟹   | 海洋     |
// +----+--------+----------+
// 2 rows in set (0.03 sec)

// mysql> select * from aas;
// +----+------+
// | id | name |
// +----+------+
// |  1 | 海洋 |
// +----+------+
// 1 row in set (0.03 sec)
func (d Dao) AssociationTags() {
	d.DB().AutoMigrate(&AAS{}, &MasterYI{})
	d.DB().Create(&AAS{
		Name: "海洋",
		MasterYI: []MasterYI{
			{Name: "石斑鱼"},
			{Name: "海蟹"},
		}},
	)
}

// 多对多
type AASB struct {
	ID        int
	Name      string
	MasterYIB []MasterYIB `gorm:"many2many:aasb_masteryib;foreignKey:Name;references:AASBName"`
}
type MasterYIB struct {
	ID       int
	Name     string
	AASBName string
}

func (d Dao) AssociationTagsMany2() {
	// d.DB().AutoMigrate(&AASB{}, &MasterYIB{})
	d.DB().Create(&AASB{
		Name: "海洋",
		MasterYIB: []MasterYIB{
			{AASBName: "石斑鱼"},
			{AASBName: "海蟹"},
			{AASBName: "鲸鱼"},
		}},
	)
}
func (d Dao) AssociationTagsMany2Select() {
	// d.DB().AutoMigrate(&AASB{}, &MasterYIB{})
	var as AASB
	d.DB().Preload("MasterYIB").Model(&AASB{}).Where("name =?", "海洋").Find(&as)
	fmt.Println(as)
}

// mysql> select * from aasb;select * from master_yib;select * from aasb_masteryib;
// +----+------+
// | id | name |
// +----+------+
// |  1 | 海洋 |
// +----+------+
// 1 row in set (0.08 sec)

// +----+------+-----------+
// | id | name | aasb_name |
// +----+------+-----------+
// |  1 |      | 石斑鱼    |
// |  2 |      | 海蟹      |
// |  3 |      | 鲸鱼      |
// +----+------+-----------+
// 3 rows in set (0.10 sec)

// +-----------+---------------------+
// | aasb_name | master_yibaasb_name |
// +-----------+---------------------+
// | 海洋      | 海蟹                |
// | 海洋      | 石斑鱼              |
// | 海洋      | 鲸鱼                |
// +-----------+---------------------+
// 3 rows in set (0.09 sec)

// joinForeignKey 标识连接表中映射回当前模型表的外键列。
// joinReferences 指向连接表中链接到参考模型表的外键列。
