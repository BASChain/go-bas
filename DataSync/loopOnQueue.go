package DataSync

import "sync"

var queryAssetLock = &sync.RWMutex{}

func loopOverQueueAsset(waitGroup *sync.WaitGroup, handle QueryHandler){
	queryAssetLock.Lock()
	defer queryAssetLock.Unlock()
	for _,s:= range queueAsset.ToSlice() {
		handle(s)
	}
	waitGroup.Done()
}
