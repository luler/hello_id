package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);comment:账号"`
	Password string `gorm:"type:varchar(100);comment:密码"`
}
