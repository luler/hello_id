package model

import (
	"time"
)

type IdRule struct {
	Id        uint   `gorm:"primarykey;autoIncrement"`
	Type      string `gorm:"type:varchar(50);not null;default:'';comment:规则标识;unique"`
	CurrentId int    `gorm:"not null;default:0;comment:当前id值"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
