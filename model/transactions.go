package model

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	TxId        string              `gorm:"index" json:"id"`
	Version     int                 `gorm:"column:version" json:"version"`
	RawTx       string              `gorm:"column:rawtx" json:"rawtx"`
	BlockNumber int64               `gorm:"index" json:"blocknumber"`
	Inputs      []TransactionInputs `gorm:"foreignKey:TxId;references:TxId" json:"inputs"`
	Outputs     []TransactionOutput `gorm:"foreignKey:TxId;references:TxId" json:"outputs"`
	LockTime    int                 `gorm:"column:locktime" json:"locktime"`
	IsSpendTx   bool                `gorm:"-" json:"is_spend_tx"`
}

type TransactionInputs struct {
	gorm.Model
	TxId         string `gorm:"index" json:"txid"`
	PrevTxId     string `gorm:"column:prev_txid" json:"prev_txid"`
	PrevOutIndex int    `gorm:"column:prev_out_index" json:"prev_out_index"`
	ScriptSig    string `gorm:"column:script_sig" json:"script_sig"`
	Witness      string `gorm:"column:witness" json:"witness"`
}

type TransactionOutput struct {
	gorm.Model
	TxId     string `gorm:"index" json:"txid"`
	Index    int    `gorm:"column:index" json:"index"`
	Value    int64  `gorm:"column:value" json:"value"`
	PkScript string `gorm:"column:pk_script" json:"pk_script"`
	Address  string `gorm:"index,index:address_utxo" json:"address"`
	Spent    bool   `gorm:"index:address_utxo" json:"spent"`
	SpentTx  string `gorm:"column:spent_tx" json:"spent_tx"`
}
