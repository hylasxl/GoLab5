package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null"`
	Age     int    `gorm:"type:int(11);not null"`
	ClassID uint
	Class   Class `gorm:"foreignKey:ClassID"`
}
