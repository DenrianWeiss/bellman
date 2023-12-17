package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var dbOnce sync.Once

func GetDb() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("chain.db"), &gorm.Config{})
		db.Set("gorm:auto_preload", true)
		if err != nil {
			panic(err)
		}
	})
	return db
}
