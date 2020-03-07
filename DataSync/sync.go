package DataSync

import (
	"fmt"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

func fillWaitQueue(lastBlockNumber uint64){
	opts := getLoopOpts(lastSavingPoint,&lastBlockNumber)

	//base data should be queried first
	loopOverMintAsset(opts,handleMintAsset)

	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	go loopOverTakeoverAsset(&waitGroup,opts,handleTakeoverAsset)
	go loopOverRechargeAsset(&waitGroup,opts,handleRechargeAsset)
	go loopOverRootChanged(&waitGroup,opts, handleRootChanged)
	go loopOverDNSRecordChange(&waitGroup,opts,handleDNSRecordChange)
	go loopOverDNSRecordRemove(&waitGroup,opts,handleDNSRecordRemove)

	waitGroup.Wait()


}

func syncDataByHandleQueue(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go loopOverQueueAsset(&waitGroup,handleAssetUpdate)
	go loopOverQueueDNS(&waitGroup, handleDNSUpdate)

	waitGroup.Wait()

	clearQueryAsset()
	clearQueueDns()

	//test code
	//for _,value:=range Records {
	//	fmt.Print(string(value.asset.Name),value.asset.Owner.String())
	//	if value.dns!=nil{
	//		fmt.Println(value.dns.Ipv4)
	//	}else{
	//		fmt.Println("")
	//	}
	//}

	a,e:=Bas_Ethereum.QueryAssetInfo(Bas_Ethereum.GetHash("root1"))
	if e==nil {
		fmt.Println("test query normal : ", string(a.Name))
	}else{
		fmt.Println("test query failed : ",e)
	}

}

var subs []event.Subscription

func watch(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup

	opts := getWatchOpts(lastBlockNumber)
	subs = []event.Subscription{}

	waitGroup.Add(6)

	go watchMintAsset(opts,&subs,&waitGroup)
	go watchTakeoverAsset(opts,&subs,&waitGroup)
	go watchRechargeAsset(opts,&subs,&waitGroup)
	go watchRootChanged(opts,&subs,&waitGroup)
	go watchDNSRecordChange(opts,&subs,&waitGroup)
	go watchDNSRecordRemove(opts,&subs,&waitGroup)

	waitGroup.Wait()

	ReSync()
}

func unSubscriptAll(){
	for _ ,s:= range subs {
		s.Unsubscribe()
	}
	logger.Info("Unsubscribed all event watches")
}

func ReSync(){
	logger.Info("ReSyncing")
	unSubscriptAll()
	Bas_Ethereum.ResetConnAndContracts()
	Sync()
}

func Sync(){
	lastBlockNumber := Bas_Ethereum.GetLastBlockNumber()
	fillWaitQueue(lastBlockNumber)
	syncDataByHandleQueue()
	lastSavingPoint = lastBlockNumber
	logger.Info("current saving point is blocknumber : ", lastSavingPoint)
	watch(lastBlockNumber)
}
