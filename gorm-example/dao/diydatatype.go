package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// scan拿出的数据，一般是[]byte
// value是原始数据，return 1是存入数据的
// https://gorm.io/zh_CN/docs/data_types.html
// https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE

// 数组切片
type Args []string

func (c Args) Value() (driver.Value, error) {
	if len(c) > 0 {
		var str string = c[0]
		for _, v := range c[1:] {
			str += "," + v
		}
		return str, nil
	}
	return "", nil
}
func (c *Args) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型无法解析")
	}

	*c = strings.Split(string(str), ",")
	return nil
}

type Info struct {
	Name     string
	Describe string
}
type DogKing struct {
	gorm.Model
	Name     string
	Age      int
	Describe Info `gorm:"type:text;size:1024;comment:归属"`
	Args     Args `gorm:"type:text"`
}

// 字符串json
func (c Info) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), err
}
func (c *Info) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型不匹配")
	}
	json.Unmarshal(str, c)
	return nil
}

func (d Dao) DataTypeString() {
	d.DB().AutoMigrate(&DogKing{}, &Info{})

	d.DB().Create(&DogKing{Describe: Info{
		Name:     "nike",
		Describe: "舔狗",
	}, Args: Args{"1", "2", "3"}})

	var dk DogKing
	d.DB().First(&dk)
	fmt.Println(dk)
}

// mysql> select * from dog_king;
// +----+---------------------+---------------------+------------+------+-----+-----------------------------------+-------+
// | id | created_at          | updated_at          | deleted_at | name | age | describe                          | args  |
// +----+---------------------+---------------------+------------+------+-----+-----------------------------------+-------+
// |  1 | 2024-06-30 19:09:04 | 2024-06-30 19:09:04 | NULL       |      |   0 | {"Name":"nike","Describe":"舔狗"} | 1,2,3 |
// +----+---------------------+---------------------+------------+------+-----+-----------------------------------+-------+
// 1 row in set (0.09 sec)
