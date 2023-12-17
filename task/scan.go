package task

import (
	"github.com/DenrianWeiss/bellman/service/rpc"
	"github.com/DenrianWeiss/bellman/service/scan"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

var url string
var auth string

func init() {
	// Read Server Url and Auth from env
	url = os.Getenv("RPC_URL")
	auth = os.Getenv("RPC_AUTH")
}

func ScanJob() {
	// Get the last block from the database
	dbLatest := scan.GetLatestBlock()
	// Get the last block from the blockchain
	chainLatest, err := rpc.GetBlockCount(url, auth)
	if err != nil {
		panic(err)
	}
	blockDiff := chainLatest - dbLatest
	// Scan the blocks
	if blockDiff > 0 {
		if blockDiff > 1000 {
			blockDiff = 1000
		}
		err := scan.ScanBlockRange(url, auth, dbLatest, dbLatest+blockDiff)
		if err != nil {
			log.Printf("Error scanning blocks: %s", err.Error())
			return
		}
	}
}

func ScanJobCron() {
	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	c.AddFunc("@every 10s", ScanJob)
	c.Start()
}
