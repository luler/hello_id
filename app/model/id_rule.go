package model

import (
	"go_test/app/helper/type_helper"
)

type IdRule struct {
	Id        uint             `gorm:"primarykey;autoIncrement;comment:自定义ID表" json:"id"`
	Type      string           `gorm:"type:varchar(50);not null;default:'';comment:规则标识;unique" json:"type"`
	CurrentId int64            `gorm:"not null;default:0;comment:当前id值" json:"currentId"`
	Step      int32            `gorm:"not null;default:0;comment:预占幅度" json:"step"`
	Prefix    string           `gorm:"type:varchar(255);not null;default:'';comment:前缀" json:"prefix"`
	Suffix    string           `gorm:"type:varchar(255);not null;default:'';comment:后缀" json:"suffix"`
	MinLength int              `gorm:"not null;default:0;comment:最小长度，如递增值长度不足，左边补0，设为0时为不限制" json:"minLength"`
	CreatedAt type_helper.Time `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt type_helper.Time `gorm:"comment:更新时间" json:"updatedAt"`
}
