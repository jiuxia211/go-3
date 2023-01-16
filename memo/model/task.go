package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKEY:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"not null"`
	Status    int    `gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
