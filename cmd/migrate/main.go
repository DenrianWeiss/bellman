package main

import (
	"github.com/DenrianWeiss/bellman/model"
	"github.com/DenrianWeiss/bellman/service/db"
	"log"
)

func main() {
	err := db.GetDb().AutoMigrate(&model.Block{}, &model.Status{}, &model.Transactions{}, &model.TransactionOutput{}, &model.TransactionInputs{}, model.SyncFail{})
	if err != nil {
		log.Printf("Error migrating database: %s", err.Error())
	}
}
