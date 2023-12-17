# Bellman

Bells Blockchain Indexer & UTXO scanner

## Installation

1. Clone this repository
2. Run `go build -o bellman .`
3. Run `go build -o migrate cmd/migrate/main.go`
4. Run `./migrate`
5. Set ENV `RPC_AUTH=<username>:<password>;RPC_URL=<bellsd_rpc>` 
6. Run `./bellman` as daemon

## Notice

1. This scanner does NOT handle reorgs. If a reorg happens, you need to reindex the whole chain.
2. This scanner does NOT handle forks. If a fork happens, you need to reindex the whole chain.

## Apis

### Get Block

GET `/block/:hash`

```json
{"block":{"ID":10,"CreatedAt":"2023-12-17T13:40:34.175877+08:00","UpdatedAt":"2023-12-17T13:40:34.175877+08:00","DeletedAt":null,"hash":"d217f6deee40f4691ca333b3bad48ce87e2c8157997d445958283c87e89afc9b","size":190,"height":10,"version":1,"merkleroot":"","time":1701898127,"nonce":18533,"bits":"1e0ffff0","difficulty":0.00024414}}
```

### Get Block by Height

GET `/block-by-num/:height`

```json
{"block":{"ID":10,"CreatedAt":"2023-12-17T13:40:34.175877+08:00","UpdatedAt":"2023-12-17T13:40:34.175877+08:00","DeletedAt":null,"hash":"d217f6deee40f4691ca333b3bad48ce87e2c8157997d445958283c87e89afc9b","size":190,"height":10,"version":1,"merkleroot":"","time":1701898127,"nonce":18533,"bits":"1e0ffff0","difficulty":0.00024414}}
```

### Get Transaction

GET `/tx/:hash`

```json
{
   "transaction":{
      "ID":17951,
      "CreatedAt":"2023-12-17T13:47:14.666424+08:00",
      "UpdatedAt":"2023-12-17T13:47:14.666424+08:00",
      "DeletedAt":null,
      "id":"94c7285fbc9b99c4070b142de4cf205a06dec91c578f9d4085a084ffdc77c394",
      "version":1,
      "rawtx":"01000000018cef66f959659fee9b508e1b096403211e556e59f39546c7de76af16c1cb9c9b010000006b4830450221008f0999ddc8dc188333bab291d5c6f7e91f040435b45c4fa6e93b078202de41f4022072416f37c02ac429c27b750065f0e8e5c5517bf80308b400dd6896dd2269b144012103f433de76d5989cf5b16d9616d1ac09116ed6dbbab3b309f7b3ff737421524b31ffffffff02d0300e00000000001976a914e1aa7b6a0c55f2660fd7c96dd11b1d06e816582e88ac409c0000000000001976a9146d533b617aa9ccad76f0a2de1fae2acf37231e2888ac00000000",
      "blocknumber":15113,
      "inputs":[
         {
            "ID":30284,
            "CreatedAt":"2023-12-17T13:47:14.666697+08:00",
            "UpdatedAt":"2023-12-17T13:47:14.666697+08:00",
            "DeletedAt":null,
            "txid":"94c7285fbc9b99c4070b142de4cf205a06dec91c578f9d4085a084ffdc77c394",
            "prev_txid":"9b9ccbc116af76dec74695f3596e551e210364091b8e509bee9f6559f966ef8c",
            "prev_out_index":1,
            "script_sig":"4830450221008f0999ddc8dc188333bab291d5c6f7e91f040435b45c4fa6e93b078202de41f4022072416f37c02ac429c27b750065f0e8e5c5517bf80308b400dd6896dd2269b144012103f433de76d5989cf5b16d9616d1ac09116ed6dbbab3b309f7b3ff737421524b31",
            "witness":""
         }
      ],
      "outputs":[
         {
            "ID":46869,
            "CreatedAt":"2023-12-17T13:47:14.666801+08:00",
            "UpdatedAt":"2023-12-17T13:47:14.666801+08:00",
            "DeletedAt":null,
            "txid":"94c7285fbc9b99c4070b142de4cf205a06dec91c578f9d4085a084ffdc77c394",
            "index":0,
            "value":930000,
            "pk_script":"76a914e1aa7b6a0c55f2660fd7c96dd11b1d06e816582e88ac",
            "address":"BR2HtBR9Pw5MZaAoYZkbF5yYQzGnxg1i1P",
            "spent":false
         },
         {
            "ID":46870,
            "CreatedAt":"2023-12-17T13:47:14.666801+08:00",
            "UpdatedAt":"2023-12-17T13:47:14.666801+08:00",
            "DeletedAt":null,
            "txid":"94c7285fbc9b99c4070b142de4cf205a06dec91c578f9d4085a084ffdc77c394",
            "index":1,
            "value":40000,
            "pk_script":"76a9146d533b617aa9ccad76f0a2de1fae2acf37231e2888ac",
            "address":"BER8ygYumndBo66qL3P338YWHb7fXtXhRt",
            "spent":false
         }
      ],
      "locktime":0
   }
}
```

### Get Transaction by Block

GET `/tx-by-num/:height`

```json
{
   "transactions":[
      {
         "ID":1446,
         "CreatedAt":"2023-12-17T13:40:47.212347+08:00",
         "UpdatedAt":"2023-12-17T13:40:47.212347+08:00",
         "DeletedAt":null,
         "id":"5855f62863feed51912f410e0fc6c057d3cc899adf25429ec37d7978805ee7ea",
         "version":1,
         "rawtx":"01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0e048de371650105062f503253482fffffffff0100f2052a0100000023210273728f16e63be15c807c6a9ec40ab66ba94eb124978b2ed2b586e480c56048ffac00000000",
         "blocknumber":1444,
         "inputs":[
            {
               "ID":1446,
               "CreatedAt":"2023-12-17T13:40:47.212992+08:00",
               "UpdatedAt":"2023-12-17T13:40:47.212992+08:00",
               "DeletedAt":null,
               "txid":"5855f62863feed51912f410e0fc6c057d3cc899adf25429ec37d7978805ee7ea",
               "prev_txid":"0000000000000000000000000000000000000000000000000000000000000000",
               "prev_out_index":4294967295,
               "script_sig":"048de371650105062f503253482f",
               "witness":""
            }
         ],
         "outputs":[
            {
               "ID":1446,
               "CreatedAt":"2023-12-17T13:40:47.213288+08:00",
               "UpdatedAt":"2023-12-17T13:40:47.213288+08:00",
               "DeletedAt":null,
               "txid":"5855f62863feed51912f410e0fc6c057d3cc899adf25429ec37d7978805ee7ea",
               "index":0,
               "value":5000000000,
               "pk_script":"210273728f16e63be15c807c6a9ec40ab66ba94eb124978b2ed2b586e480c56048ffac",
               "address":"0273728f16e63be15c807c6a9ec40ab66ba94eb124978b2ed2b586e480c56048ff",
               "spent":false
            }
         ],
         "locktime":0
      }
   ]
}
```

### Get Txs by Address

GET `/txs/:address`

```json
```