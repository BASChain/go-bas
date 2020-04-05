package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
)

var queueOwnership = new([]Bas_Ethereum.Hash)
var queueRoot = new([]Bas_Ethereum.Hash)
var queueSub = new([]Bas_Ethereum.Hash)
var queueDns = new([]Bas_Ethereum.Hash)

var firstAppearInBlock = make(map[Bas_Ethereum.Hash]uint64)

var oLock = &sync.RWMutex{}
var rLock = &sync.RWMutex{}
var sLock = &sync.RWMutex{}
var dLock = &sync.RWMutex{}


func insertQueue(key Bas_Ethereum.Hash,queue *[]Bas_Ethereum.Hash){
	if !exists(*queue,key) {
		*queue = append(*queue, key)
	}
}

func clearQueueO(){
	oLock.Lock()
	defer oLock.Unlock()
	queueOwnership = new([]Bas_Ethereum.Hash)
}

func clearQueueR(){
	rLock.Lock()
	defer rLock.Unlock()
	queueRoot = new([]Bas_Ethereum.Hash)
}

func clearQueueS(){
	sLock.Lock()
	defer sLock.Unlock()
	queueSub = new([]Bas_Ethereum.Hash)
}

func clearQueueD(){
	dLock.Lock()
	defer dLock.Unlock()
	queueDns = new([]Bas_Ethereum.Hash)
}


func loopOverQueueOwnership(waitGroup *sync.WaitGroup){
	oLock.Lock()
	defer oLock.Unlock()
	for _,s:= range *queueOwnership {
		go updateByQueryOwnership(s, 0)
	}
	waitGroup.Done()
}

func loopOverQueueRoot(waitGroup *sync.WaitGroup){
	rLock.Lock()
	defer rLock.Unlock()
	for _,s:= range *queueRoot {
		go updateByQueryRoot(s,0)
	}
	waitGroup.Done()
}

func loopOverQueueSub(waitGroup *sync.WaitGroup){
	sLock.Lock()
	defer sLock.Unlock()
	for _,s:= range *queueSub {
		go updateByQuerySub(s,0)
	}
	waitGroup.Done()
}

func loopOverQueueDns(waitGroup *sync.WaitGroup){
	dLock.Lock()
	defer dLock.Unlock()
	for _,s:= range *queueDns {
		go updateByQueryDNS(s,0)
	}
	waitGroup.Done()
}
