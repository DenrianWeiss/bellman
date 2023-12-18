package web

import "github.com/gin-gonic/gin"

func RegisterRoute(r *gin.Engine) {
	r.GET("/latest-block", HandleGetLatestBlock)
	r.GET("/tx/:hash", HandleGetTxByHash)
	r.GET("/block/:hash", HandleGetBlockByHash)
	r.GET("/txs-by-num/:blockNumber", HandleGetTxByBlockNumber)
	r.GET("/block-by-num/:blockNumber", HandleBlockByNumber)
	r.GET("/txs/:address", HandleGetTxByAddress)
	r.GET("/recent-txs/:address", HandleGetRecentTxs)
	r.GET("/utxo/:address", HandleGetUtxoByAddress)
	r.GET("/ping", HandlePing)
	r.Any("/rpc", ProxyJsonRpc)
}
