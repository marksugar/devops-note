package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	Name       string      `json:"name"`
	Code       string      `json:"code"`
	Operations []Operation `gorm:"many2many:role_operations;"`
}

type Operation struct {
	gorm.Model
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
	RoleID uint   `json:"role_id"`
}
