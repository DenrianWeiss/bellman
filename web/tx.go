package web

import (
	"github.com/DenrianWeiss/bellman/model"
	"github.com/DenrianWeiss/bellman/service/db"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetTxByHash(hash string) (model.Transactions, error) {
	var tx model.Transactions
	err := db.GetDb().Where("tx_id = ?", hash).First(&tx).Error
	// Get Related UTXOs
	var inputs []model.TransactionInputs
	var outputs []model.TransactionOutput
	err = db.GetDb().Where("tx_id = ?", hash).Find(&inputs).Error
	if err != nil {
		return tx, err
	}
	err = db.GetDb().Where("tx_id = ?", hash).Find(&outputs).Error
	if err != nil {
		return tx, err
	}
	tx.Inputs = inputs
	tx.Outputs = outputs
	return tx, err
}

func GetBlockByHash(hash string) (model.Block, error) {
	var block model.Block
	err := db.GetDb().Where("hash = ?", hash).First(&block).Error
	return block, err
}

func GetBlockByNumber(blockNumber int64) (model.Block, error) {
	var block model.Block
	err := db.GetDb().Where("height = ?", blockNumber).First(&block).Error
	return block, err
}

func GetTxByBlockNumber(blockNumber int64) ([]model.Transactions, error) {
	var txs []model.Transactions
	err := db.GetDb().Where("block_number = ?", blockNumber).Find(&txs).Error
	// Add inputs and outputs to txs
	for i, tx := range txs {
		var inputs []model.TransactionInputs
		var outputs []model.TransactionOutput
		err = db.GetDb().Where("tx_id = ?", tx.TxId).Find(&inputs).Error
		if err != nil {
			return nil, err
		}
		err = db.GetDb().Where("tx_id = ?", tx.TxId).Find(&outputs).Error
		if err != nil {
			return nil, err
		}
		txs[i].Inputs = inputs
		txs[i].Outputs = outputs
	}
	return txs, err
}

func GetTxByAddress(address string) ([]model.Transactions, error) {
	var txInputs []model.TransactionOutput
	var txOutputs []model.TransactionOutput
	var txs []model.Transactions
	err := db.GetDb().Where("address = ?", address).Find(&txInputs).Error
	if err != nil {
		return nil, err
	}
	err = db.GetDb().Where("address = ?", address).Find(&txOutputs).Error
	if err != nil {
		return nil, err
	}
	for _, txInput := range txInputs {
		var tx model.Transactions
		err = db.GetDb().Where("tx_id = ?", txInput.TxId).First(&tx).Error
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	for _, txOutput := range txOutputs {
		var tx model.Transactions
		err = db.GetDb().Where("tx_id = ?", txOutput.TxId).First(&tx).Error
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	// Add inputs and outputs to txs
	for i, tx := range txs {
		var inputs []model.TransactionInputs
		var outputs []model.TransactionOutput
		err = db.GetDb().Where("tx_id = ?", tx.TxId).Find(&inputs).Error
		if err != nil {
			return nil, err
		}
		err = db.GetDb().Where("tx_id = ?", tx.TxId).Find(&outputs).Error
		if err != nil {
			return nil, err
		}
		txs[i].Inputs = inputs
		txs[i].Outputs = outputs
	}
	return txs, err
}

func HandleGetLatestBlock(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"block": GetLatestBlock()})
}

func HandleGetTxByHash(ctx *gin.Context) {
	hash := ctx.Param("hash")
	tx, err := GetTxByHash(hash)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"transaction": tx})
}

func HandleGetBlockByHash(ctx *gin.Context) {
	hash := ctx.Param("hash")
	block, err := GetBlockByHash(hash)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"block": block})
}

func HandleBlockByNumber(ctx *gin.Context) {
	num := ctx.Param("blockNumber")
	// Convert num to int64
	numInt, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	block, err := GetBlockByNumber(numInt)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"block": block})
}

func HandleGetTxByBlockNumber(ctx *gin.Context) {
	num := ctx.Param("blockNumber")
	// Convert num to int64
	numInt, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	txs, err := GetTxByBlockNumber(numInt)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"transactions": txs})
}

func HandleGetTxByAddress(ctx *gin.Context) {
	address := ctx.Param("address")
	txs, err := GetTxByAddress(address)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"transactions": txs})
}
