package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null"`
	Students []Student `gorm:"foreignKey:ClassID"`
}
