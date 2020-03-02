package DataSync

import (
	"fmt"
	"github.com/BASChain/go-bas/Bas_Ethereum"
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
	for _,value:=range Records {
		fmt.Print(string(value.asset.Name),value.asset.Owner.String())
		if value.dns!=nil{
			fmt.Println(value.dns.Ipv4)
		}else{
			fmt.Println("")
		}
	}
}

func watch(lastBlockNumber uint64){
	opts := getWatchOpts(lastBlockNumber)
	go watchMintAsset(opts)
	go watchTakeoverAsset(opts)
	go watchRechargeAsset(opts)
	go watchRootChanged(opts)
	go watchDNSRecordChange(opts)
	go watchDNSRecordRemove(opts)
}

func Sync(){
	lastBlockNumber := Bas_Ethereum.GetLastBlockNumber()

	fillWaitQueue(lastBlockNumber)
	syncDataByHandleQueue()
	lastSavingPoint = lastBlockNumber

	watch(lastBlockNumber)


}
