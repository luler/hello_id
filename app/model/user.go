package model

import "time"

type User struct {
	Id        uint   `gorm:"primarykey;autoIncrement"`
	Name      string `gorm:"type:varchar(50);default:'';comment:账号;unique"`
	Password  string `gorm:"type:varchar(100);default:'';comment:密码"`
	Status    int8   `gorm:"default:0;comment:状态，0-禁用，1-启用"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
