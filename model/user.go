package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);default:'';comment:账号;unique"`
	Password string `gorm:"type:varchar(100);default:'';comment:密码"`
	Status   int8   `gorm:"default:0;comment:状态，0-禁用，1-启用"`
}
