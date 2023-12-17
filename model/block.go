package model

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	Hash       string  `gorm:"index" json:"hash"`
	Size       int     `gorm:"column:size" json:"size"`
	Height     int     `gorm:"index" json:"height"`
	Version    int     `gorm:"column:version" json:"version"`
	Merkleroot string  `gorm:"column:merkleroot" json:"merkleroot"`
	Time       int     `gorm:"column:time" json:"time"`
	Nonce      int     `gorm:"column:nonce" json:"nonce"`
	Bits       string  `gorm:"column:bits" json:"bits"`
	Difficulty float64 `gorm:"column:difficulty" json:"difficulty"`
}
