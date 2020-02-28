package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
)

type EventHandler func(iterator interface{})
type QueryHandler func(key interface{})

func handleMintAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetMintAssetIterator)
	insertWaitQueueAsset(string(it.Event.Name))
}

func handleAssetUpdate(key interface{}){
	k:=key.(string)
	record,err:=Bas_Ethereum.QueryAssetInfo(k)
	if err==nil{
		Records[k] = DomainRecord{
			asset: &record,
			dns:Records[k].dns,
		}
	}else{
		logger.Error("query asset of key : " + k + " error : " , err )
	}
}

