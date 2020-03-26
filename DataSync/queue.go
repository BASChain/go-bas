package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
)





var queueOwnership = new([]Bas_Ethereum.Hash)
var queueRoot = new([]Bas_Ethereum.Hash)
var queueSub = new([]Bas_Ethereum.Hash)
var queueDns = new([]Bas_Ethereum.Hash)

var fistAppearInBlock = make(map[Bas_Ethereum.Hash]uint64)

var oLock = &sync.RWMutex{}
var rLock = &sync.RWMutex{}
var sLock = &sync.RWMutex{}
var dLock = &sync.RWMutex{}

func insertQueueOwnership(key Bas_Ethereum.Hash,blockNumber uint64){
	if !exists(*queueOwnership,key) {
		*queueOwnership = append(*queueOwnership, key)
		fistAppearInBlock[key] = blockNumber
	}else{
		if blockNumber < fistAppearInBlock[key] {
			fistAppearInBlock[key] = blockNumber
		}
	}
}

func insertQueueOther(key Bas_Ethereum.Hash,queue *[]Bas_Ethereum.Hash){
	if !exists(*queue,key) {
		*queue = append(*queue, key)
	}
}

func clearQuery(queue *[]Bas_Ethereum.Hash){
	queue = new([]Bas_Ethereum.Hash)
}


func loopOverQueueOwnership(waitGroup *sync.WaitGroup){
	oLock.Lock()
	defer oLock.Unlock()
	for _,s:= range *queueOwnership {
		updateByQueryOwnership(s,fistAppearInBlock[s])
	}
	waitGroup.Done()
}

func loopOverQueueRoot(waitGroup *sync.WaitGroup){
	rLock.Lock()
	defer rLock.Unlock()
	for _,s:= range *queueRoot {
		updateByQueryRoot(s,0)
	}
	waitGroup.Done()
}

func loopOverQueueSub(waitGroup *sync.WaitGroup){
	sLock.Lock()
	defer sLock.Unlock()
	for _,s:= range *queueSub {
		updateByQuerySub(s,0)
	}
	waitGroup.Done()
}

func loopOverQueueDns(waitGroup *sync.WaitGroup){
	dLock.Lock()
	defer dLock.Unlock()
	for _,s:= range *queueDns {
		updateByQueryDNS(s,0)
	}
	waitGroup.Done()
}
