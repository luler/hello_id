package model

import (
	"time"
)

type AuthKey struct {
	Id        uint   `gorm:"primarykey;autoIncrement"`
	AuthKey   string `gorm:"type:varchar(50);default:'';comment:授权码;unique"`
	Remark    string `gorm:"type:varchar(255);default:'';comment:用途描述"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
