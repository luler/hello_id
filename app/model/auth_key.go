package model

import (
	"time"
)

type AuthKey struct {
	Id        uint      `gorm:"primarykey;autoIncrement" json:"id"`
	AuthKey   string    `gorm:"type:varchar(50);not null;default:'';comment:授权码;unique" json:"authKey"`
	Remark    string    `gorm:"type:varchar(255);not null;default:'';comment:用途描述" json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
