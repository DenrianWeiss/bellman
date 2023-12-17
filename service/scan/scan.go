package scan

import (
	"github.com/DenrianWeiss/bellman/constants"
	"github.com/DenrianWeiss/bellman/model"
	"github.com/DenrianWeiss/bellman/service/db"
	"github.com/DenrianWeiss/bellman/service/decode"
	"github.com/DenrianWeiss/bellman/service/rpc"
	"gorm.io/gorm"
	"log"
)

func GetLatestBlock() int64 {
	var status model.Status
	db.GetDb().First(&status)
	if status.LastBlock == 0 {
		// If not found, return 1
		return 1
	}
	return status.LastBlock
}

func UpdateLatestBlock(id int64) {
	// Get last block, if not found, create it
	var status model.Status
	db.GetDb().First(&status)
	if status.LastBlock == 0 {
		db.GetDb().Create(&model.Status{LastBlock: id})
		return
	}
	// If found, update it
	db.GetDb().Model(&model.Status{}).Where("id = ?", status.ID).Update("last_block", id)
}

func RecordFailedTxOrBlock(hash, t, reason string) {
	db.GetDb().Create(&model.SyncFail{Type: t, Hash: hash})
}

func UpdateBlockInDb(block model.Block) error {
	return db.GetDb().Transaction(func(tx *gorm.DB) error {
		// First check if the block exists
		var existingBlock model.Block
		tx.Where("hash = ?", block.Hash).First(&existingBlock)
		if existingBlock.Hash != "" {
			// If it exists, end tx
			return nil
		}
		// If it doesn't exist, create it
		err := tx.Create(&block).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func UpdateTxInDb(blockTx model.Transactions) error {
	db.GetDb().Transaction(func(tx *gorm.DB) error {
		// First check if the tx exists
		var existingTx model.Transactions
		tx.Where("tx_id = ?", blockTx.TxId).First(&existingTx)
		if existingTx.TxId != "" {
			// If it exists, end tx
			return nil
		}
		// If it doesn't exist, create it
		err := tx.Create(&blockTx).Error
		if err != nil {
			return err
		}
		// Update spent outputs
		for _, input := range blockTx.Inputs {
			if input.PrevTxId == constants.CoinBaseHash {
				continue
			}
			err = tx.Model(&model.TransactionOutput{}).Where("tx_id = ? AND `index` = ?", input.PrevTxId, input.PrevOutIndex).Update("spent", true).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func ScanBlockRange(url, auth string, start, end int64) error {
	// Get Every Block Using RPC
	for i := start; i <= end; i++ {
		log.Printf("Scanning block %d", i)
		hash, err := rpc.GetBlockHash(url, int(i), auth)
		if err != nil {
			return err
		}
		block, err := rpc.GetBlockByHash(url, hash, auth)
		if err != nil {
			RecordFailedTxOrBlock(hash, "block", err.Error())
			continue
		}
		// Save Block to DB
		blk, err := rpc.BlockToDbModel(block)
		if err != nil {
			return err
		}
		err = UpdateBlockInDb(*blk)
		if err != nil {
			return err
		}
		// Scan Txs
		for _, tx := range block.Result.Tx {
			// Get Raw Tx
			rawTx, err := rpc.GetRawTransaction(url, tx, auth)
			if err != nil {
				log.Printf("Error getting raw tx %s: %s", tx, err.Error())
				RecordFailedTxOrBlock(tx, "tx", err.Error())
				continue
			}
			// Decode Tx
			decoded, err := decode.TxToDbModel(rawTx, i)
			if err != nil {
				return err
			}
			// Save Tx to DB
			err = UpdateTxInDb(decoded)
			if err != nil {
				return err
			}
			UpdateLatestBlock(i)
		}
	}
	return nil
}
