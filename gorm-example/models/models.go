package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/models.html#%E6%A8%A1%E5%9E%8B%E5%AE%9A%E4%B9%89
type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

// https://gorm.io/zh_CN/docs/models.html#gorm-Model
type UserTwo struct {
	gorm.Model                  // https://gorm.io/zh_CN/docs/models.html#gorm-Model
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
}

// 主键标签，字段标签详细说明
// see https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE
type UserId struct {
	UUID uint `gorm:"primaryKey"`
	Time *time.Time
}
type UserInfo struct {
	UserId    UserId `gorm:"embedded"`
	Name      string
	Age       int `gorm:"default:18"` // 默认值18
	Birthday  int `gorm:"comment:年龄"` // comment
	Phone     int // `gorm:"not null"`   // 不能为空
	CreatedAt time.Time
	UpdatedAt time.Time
}
