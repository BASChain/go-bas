package DataSync

import "sync"
import "github.com/deckarep/golang-set"


var lockQueueAsset = &sync.RWMutex{}

var queueAsset = mapset.NewSet()

func insertWaitQueueAsset(key string){
	lockQueueAsset.Lock()
	defer lockQueueAsset.Unlock()
	queueAsset.Add(key)
}
