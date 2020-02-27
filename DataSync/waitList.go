package DataSync

import "sync"
import "github.com/deckarep/golang-set"


var waitQueueLock = &sync.RWMutex{}

var waitQueue = mapset.NewSet()

func insertWaitQueue(key string){
	waitQueueLock.Lock()
	defer waitQueueLock.Unlock()
	waitQueue.Add(key)
}



