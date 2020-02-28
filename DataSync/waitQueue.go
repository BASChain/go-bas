package DataSync

import "sync"
import "github.com/deckarep/golang-set"


var waitQueueAssetLock = &sync.RWMutex{}

var waitQueueAsset = mapset.NewSet()

func insertWaitQueueAsset(key string){
	waitQueueAssetLock.Lock()
	defer waitQueueAssetLock.Unlock()
	waitQueueAsset.Add(key)
}
