package web

import (
	"bytes"
	"encoding/base64"
	"github.com/DenrianWeiss/bellman/task"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

var AllowedMethods = map[string]bool{
	"getblock":              true,
	"getbestblockhash":      true,
	"getblockchaininfo":     true,
	"getblockcount":         true,
	"getblockfilter":        true,
	"getblockhash":          true,
	"getblockheader":        true,
	"getblockstats":         true,
	"getchaintips":          true,
	"getchaintxstats":       true,
	"getdifficulty":         true,
	"getmempoolancestors":   true,
	"getmempooldescendants": true,
	"getmempoolentry":       true,
	"getmempoolinfo":        true,
	"getrawmempool":         true,
	"gettxout":              true,
	"gettxoutproof":         true,
	"gettxoutsetinfo":       true,
	"preciousblock":         true,
	"pruneblockchain":       true,
	"savemempool":           true,
	"scantxoutset":          true,
	"verifychain":           true,
	"verifytxoutproof":      true,
	"getpeerinfo":           true,
	"createrawtransaction":  true,
	"decoderawtransaction":  true,
	"decodescript":          true,
	"getrawtransaction":     true,
	"sendrawtransaction":    true,
	"testmempoolaccept":     true,
	"validateaddress":       true,
	"verifymessage":         true,
}

type JsonRpcRequest struct {
	JsonRpc string      `json:"jsonrpc" binding:"required"`
	Method  string      `json:"method" binding:"required"`
	Params  interface{} `json:"params"`
	Id      interface{} `json:"id" binding:"required"`
}

func ProxyJsonRpc(ctx *gin.Context) {
	// Decode JSONRPC 1.0 post body
	var jsonRpcRequest JsonRpcRequest
	err := ctx.BindJSON(&jsonRpcRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Check if method is allowed
	if !AllowedMethods[jsonRpcRequest.Method] {
		ctx.JSON(400, gin.H{"error": "Method not allowed"})
		return
	}
	// Send request to bitcoind
	// Assemble Req
	jsonBody, err := json.Marshal(&jsonRpcRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req, err := http.NewRequest("POST", task.GetUrl(), bytes.NewBuffer(jsonBody))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(task.GetAuth())))
	// Send Req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	// Forward response
	ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	return
}
