package model

import "time"

type User struct {
	Id        uint      `gorm:"primarykey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null;default:'';comment:账号;unique" json:"name"`
	Password  string    `gorm:"type:varchar(100);not null;default:'';comment:密码" json:"password"`
	Status    int8      `gorm:"not null;default:0;comment:状态，0-禁用，1-启用" json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
