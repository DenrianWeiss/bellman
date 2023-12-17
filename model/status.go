package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	LastBlock int64 `gorm:"column:last_block" json:"last_block"`
}
