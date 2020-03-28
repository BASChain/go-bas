package service

import (
	"github.com/BASChain/go-bas/DataSync"
	"github.com/BASChain/go-bas/Market"
)

func StartService() {
	go DataSync.Sync()
	Market.Sync()

}