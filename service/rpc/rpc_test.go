package rpc

import "testing"

const RpcUrl = "https://bel-rpc.fubuki.app/"
const Auth = "test:test"

func TestGetBlockCount(t *testing.T) {
	blockCount, err := GetBlockCount(RpcUrl, Auth)
	if err != nil {
		t.Error(err)
	}
	if blockCount < 0 {
		t.Error("Block count is negative")
	}
}

func TestGetBlockHash(t *testing.T) {
	blockHash, err := GetBlockHash(RpcUrl, 0, Auth)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Block hash: %s", blockHash)
}

func TestGetBlockByHash(t *testing.T) {
	blockHash, err := GetBlockHash(RpcUrl, 0, Auth)
	if err != nil {
		t.Error(err)
	}
	block, err := GetBlockByHash(RpcUrl, blockHash, Auth)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Block hash: %s", block.Result.Hash)
	t.Logf("Block height: %d", block.Result.Height)
}
