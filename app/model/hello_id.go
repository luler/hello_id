package model

import (
	"time"
)

type HelloId struct {
	Id        uint   `gorm:"primarykey;autoIncrement"`
	Type      string `gorm:"type:varchar(50);default:'';comment:标识;unique"`
	CurrentId int    `gorm:"default:0;comment:当前id值"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
