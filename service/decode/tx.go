package decode

import (
	"encoding/hex"
	"errors"
	"github.com/DenrianWeiss/bellman/model"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log"
	"strings"
)

var errNonStandardOutput = errors.New("non-standard or currently unsupported script")

func DecodeTx(rawTx string) (*btcutil.Tx, error) {
	// First decode the raw tx from hex
	txBytes, err := hex.DecodeString(strings.TrimPrefix(rawTx, "0x"))
	if err != nil {
		return nil, err
	}
	tx, err := btcutil.NewTxFromBytes(txBytes)
	if err != nil {
		return nil, err
	}
	tx.MsgTx()
	return tx, nil
}

func DecodedOutPutToAddress(decodedOutput wire.TxOut) (string, error) {
	//scriptClass, err := txscript.ParsePkScript(decodedOutput.PkScript)
	_, address, _, err := txscript.ExtractPkScriptAddrs(decodedOutput.PkScript, GetChainParams())
	if err != nil {
		return "", err
	}
	if len(address) == 0 {
		log.Printf("non-standard or currently unsupported script: %s", hex.EncodeToString(decodedOutput.PkScript))
		return "", errNonStandardOutput
	}
	return address[0].String(), nil
}

func TxToDbModel(rawTx string, blockNumber int64) (model.Transactions, error) {
	tx, err := DecodeTx(rawTx)
	if err != nil {
		return model.Transactions{}, err
	}
	var inputs []model.TransactionInputs
	var outputs []model.TransactionOutput
	for _, input := range tx.MsgTx().TxIn {
		inputs = append(inputs, model.TransactionInputs{
			TxId:         tx.Hash().String(),
			PrevTxId:     input.PreviousOutPoint.Hash.String(),
			PrevOutIndex: int(input.PreviousOutPoint.Index),
			ScriptSig:    hex.EncodeToString(input.SignatureScript),
		})
	}
	for i, output := range tx.MsgTx().TxOut {
		address, err := DecodedOutPutToAddress(*output)
		if errors.Is(err, errNonStandardOutput) {
			continue
		}
		if err != nil {
			return model.Transactions{}, err
		}
		outputs = append(outputs, model.TransactionOutput{
			TxId:     tx.Hash().String(),
			Index:    i,
			Value:    output.Value,
			PkScript: hex.EncodeToString(output.PkScript),
			Address:  address,
			Spent:    false,
		})
	}
	return model.Transactions{
		TxId:        tx.Hash().String(),
		Version:     int(tx.MsgTx().Version),
		RawTx:       rawTx,
		BlockNumber: blockNumber,
		Inputs:      inputs,
		Outputs:     outputs,
		LockTime:    int(tx.MsgTx().LockTime),
	}, nil
}
