package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	User      User `gorm:"ForeignKEY:Uid"`
	Uid       uint `gorm:"not null"`
	Title     string
	Content   string `gorm:"type:longtext"`
	Status    int    `gorm:"default:'0'"` //0为未完成1为已完成
	StartTime int64
	EndTime   int64
}
