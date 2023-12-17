package model

import "gorm.io/gorm"

type SyncFail struct {
	gorm.Model
	Type   string `gorm:"column:type" json:"type"`
	Hash   string `gorm:"column:hash" json:"hash"`
	Reason string `gorm:"column:reason" json:"reason"`
}
