package rpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/bellman/model"
	"io"
	"log"
	"net/http"
)

type JsonRpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

type GetBlockHashResponse struct {
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
}

type GetBlockResponse struct {
	Result struct {
		Hash          string   `json:"hash"`
		Confirmations int      `json:"confirmations"`
		Size          int      `json:"size"`
		Height        int      `json:"height"`
		Version       int      `json:"version"`
		Merkleroot    string   `json:"merkleroot"`
		Tx            []string `json:"tx"`
		Time          int      `json:"time"`
		Nonce         int      `json:"nonce"`
		Bits          string   `json:"bits"`
		Difficulty    float64  `json:"difficulty"`
		Nextblockhash string   `json:"nextblockhash"`
	} `json:"result"`
	Error interface{} `json:"error"`
	Id    int         `json:"id"`
}

type GetRawTransactionResponse struct {
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
}

type GetBlockCountResponse struct {
	Result int         `json:"result"`
	Error  interface{} `json:"error"`
	Id     int         `json:"id"`
}

func GetBlockHash(url string, blockHeight int, basicAuth string) (string, error) {
	request := JsonRpcRequest{
		Jsonrpc: "1.0",
		Method:  "getblockhash",
		Params:  []int{blockHeight},
		Id:      1,
	}

	// Serialize and send the request
	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	// Add basic auth header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(basicAuth))))

	var response GetBlockHashResponse
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer body.Body.Close()
	resp := json.NewDecoder(body.Body)
	err = resp.Decode(&response)
	if err != nil {
		return "", err
	}
	if response.Error != nil {
		return "", fmt.Errorf("%v", response.Error)
	}

	return response.Result, nil
}

func GetBlockByHash(url string, blockHash string, basicAuth string) (*GetBlockResponse, error) {
	request := JsonRpcRequest{
		Jsonrpc: "1.0",
		Method:  "getblock",
		Params:  []string{blockHash},
		Id:      1,
	}

	// Serialize and send the request
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	// Add basic auth header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(basicAuth))))

	response := &GetBlockResponse{}
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer body.Body.Close()
	bodyPayload, _ := io.ReadAll(body.Body)
	log.Printf("Body: %s", bodyPayload)
	err = json.Unmarshal(bodyPayload, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func BlockToDbModel(block *GetBlockResponse) (*model.Block, error) {
	return &model.Block{
		Hash:       block.Result.Hash,
		Size:       block.Result.Size,
		Height:     block.Result.Height,
		Version:    block.Result.Version,
		Time:       block.Result.Time,
		Nonce:      block.Result.Nonce,
		Bits:       block.Result.Bits,
		Difficulty: block.Result.Difficulty,
	}, nil
}

func GetRawTransaction(url string, txHash string, basicAuth string) (string, error) {
	request := JsonRpcRequest{
		Jsonrpc: "1.0",
		Method:  "getrawtransaction",
		Params:  []interface{}{txHash},
		Id:      1,
	}

	// Serialize and send the request
	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	// Add basic auth header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(basicAuth))))

	var response GetRawTransactionResponse
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bodyPayload, _ := io.ReadAll(body.Body)
	log.Printf("Body: %s", bodyPayload)
	err = json.Unmarshal(bodyPayload, &response)
	if err != nil {
		return "", err
	}
	if response.Error != nil {
		return "", fmt.Errorf("%v", response.Error)
	}

	return response.Result, nil
}

func GetBlockCount(url string, basicAuth string) (int64, error) {
	request := JsonRpcRequest{
		Jsonrpc: "1.0",
		Method:  "getblockcount",
		Params:  []interface{}{},
		Id:      1,
	}

	// Serialize and send the request
	reqBody, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return 0, err
	}
	// Add basic auth header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(basicAuth))))

	var response GetBlockCountResponse
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer body.Body.Close()
	resp := json.NewDecoder(body.Body)
	err = resp.Decode(&response)
	if err != nil {
		return 0, err
	}
	if response.Error != nil {
		return 0, fmt.Errorf("%v", response.Error)
	}

	return int64(response.Result), nil
}
