package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
)


var queryAssetLock = &sync.RWMutex{}
var queryDNSLock = &sync.RWMutex{}

type QueryHandler func(key interface{})

func handleAssetUpdate(key interface{}){
	updateAsset(key.(Bas_Ethereum.Hash))
}

func handleDNSUpdate(key interface{}){
	updateDNS(key.(Bas_Ethereum.Hash))
}

func loopOverQueueAsset(waitGroup *sync.WaitGroup, handle QueryHandler){
	queryAssetLock.Lock()
	defer queryAssetLock.Unlock()
	for _,s:= range queueAsset {
		handle(s)
	}
	waitGroup.Done()
}

func loopOverQueueDNS(waitGroup *sync.WaitGroup, handle QueryHandler){
	queryDNSLock.Lock()
	defer queryDNSLock.Unlock()
	for _,s:=range queueDns {
		handle(s)
	}
	waitGroup.Done()
}