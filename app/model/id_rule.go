package model

import (
	"time"
)

type IdRule struct {
	Id        uint   `gorm:"primarykey;autoIncrement"`
	Type      string `gorm:"type:varchar(50);not null;default:'';comment:规则标识;unique"`
	CurrentId int64  `gorm:"not null;default:0;comment:当前id值"`
	Prefix    string `gorm:"type:varchar(255);not null;default:'';comment:前缀"`
	Suffix    string `gorm:"type:varchar(255);not null;default:'';comment:后缀"`
	MinLength int    `gorm:"not null;default:0;comment:最小长度，如递增值长度不足，左边补0，设为0时为不限制"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
