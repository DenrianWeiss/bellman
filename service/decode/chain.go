package decode

import (
	"github.com/DenrianWeiss/bellman/constants"
	"github.com/btcsuite/btcd/chaincfg"
)

func GetChainParams() *chaincfg.Params {
	return &chaincfg.Params{
		PubKeyHashAddrID: constants.AddressPrefix,
		ScriptHashAddrID: constants.ScriptPrefix,
		PrivateKeyID:     constants.WifPrefix,
	}
}
