package decode

import (
	"github.com/DenrianWeiss/bellman/constants"
	"testing"
)

func TestDecodeTx(t *testing.T) {
	rawTx := "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0b020d2c5993fc01a5050000000000000100743ba40b0000001976a9140a1e53de6d6a77e5da5de21abc34350f18c581a188ac00000000"
	tx, err := DecodeTx(rawTx)
	if err != nil {
		t.Error(err)
	}
	if tx.Hash().String() != "f1f047588bb418391685224436a53f1586e38e5a5c4bc8f71402ced2909bfe52" {
		t.Error("Wrong hash")
	}
	msg := tx.MsgTx()
	if len(msg.TxIn) != 1 {
		t.Error("Wrong number of inputs")
	}
	if msg.TxIn[0].PreviousOutPoint.Hash.String() != constants.CoinBaseHash {
		t.Error("Wrong previous outpoint hash")
	}
	address, err := DecodedOutPutToAddress(*tx.MsgTx().TxOut[0])
	if err != nil {
		t.Error(err)
	}
	if address != "B5NahqHQdJKxf9EkTyEuLKszYam3ywaVsy" {
		t.Error("Wrong address")
	}
}

func TestDecodeTx2(t *testing.T) {
	tx := "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0e0478e570650101062f503253482fffffffff0100c2eb0b00000000232103ba86c9bd5aabbed71342da45b4f599ec7dc3ec487d377e73766cd9a7e647a1f5ac00000000"
	dtx, err := DecodeTx(tx)
	if err != nil {
		t.Error(err)
	}
	address, err := DecodedOutPutToAddress(*dtx.MsgTx().TxOut[0])
	if err != nil {
		t.Error(err)
	}
	t.Logf("Address: %s", address)
}
