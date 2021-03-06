package service

import (
	"github.com/BASChain/go-bas/DataSync"
	"github.com/BASChain/go-bas/Market"
	"github.com/BASChain/go-bas/Miner"
	"github.com/BASChain/go-bas/utils"
)

func StartService() {
	go DataSync.Sync()
	go Market.Sync()

	go Miner.Sync()

	utils.GetBlockNum2Time().Run()

}

func StopService()  {
	utils.GetBlockNum2Time().Stop()
}