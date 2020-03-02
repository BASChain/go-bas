package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
)

var queueAsset []Bas_Ethereum.Hash
var queueDns []Bas_Ethereum.Hash

var queryAssetLock = &sync.RWMutex{}
var queryDNSLock = &sync.RWMutex{}


func insertQueueAsset(key Bas_Ethereum.Hash){
	if !exists(queueAsset,key) {
		queueAsset = append(queueAsset, key)
	}
}


func clearQueryAsset(){
	queueAsset = []Bas_Ethereum.Hash{}
}

func insertQueueDns(key Bas_Ethereum.Hash){
	if !exists(queueDns,key){
		queueDns = append(queueDns, key)
	}
}

func clearQueueDns(){
	queueDns = []Bas_Ethereum.Hash{}
}



func loopOverQueueAsset(waitGroup *sync.WaitGroup, handle updateHandler){
	queryAssetLock.Lock()
	defer queryAssetLock.Unlock()
	for _,s:= range queueAsset {
		handle(s)
	}
	waitGroup.Done()
}

func loopOverQueueDNS(waitGroup *sync.WaitGroup, handle updateHandler){
	queryDNSLock.Lock()
	defer queryDNSLock.Unlock()
	for _,s:=range queueDns {
		handle(s)
	}
	waitGroup.Done()
}