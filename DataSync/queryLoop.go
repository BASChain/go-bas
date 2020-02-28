package DataSync

import "sync"
import "github.com/BASChain/go-bas/Bas_Ethereum"

var queryAssetLock = &sync.RWMutex{}

func loopOverQueueAsset(waitGroup *sync.WaitGroup){
	queryAssetLock.Lock()
	defer queryAssetLock.Unlock()
	for _,s:= range waitQueueAsset.ToSlice() {
		key:=s.(string)
		logger.Info("querying asset of key : " + key)
		record,err:=Bas_Ethereum.QueryAssetInfo(key)
		if err==nil{
			Records[key] = DomainRecord{
				asset: &record,
				dns:Records[key].dns,
			}
		}else{
			logger.Error("query asset of key : " + key + " error : " ,err )
		}
	}
	waitGroup.Done()
}
