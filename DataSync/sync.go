package DataSync

import (
	"fmt"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"strconv"
	"sync"
)

func fillWaitQueue(){
	lastBlockNumber := Bas_Ethereum.GetLastBlockNumber()
	opts := getLoopOpts(lastSavingPoint,&lastBlockNumber)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go loopOverMintAsset(&waitGroup,opts,handleMintAsset)


	waitGroup.Wait()

	lastSavingPoint = lastBlockNumber
	logger.Info("Collected waitQueue to block number: " + strconv.FormatUint(lastSavingPoint, 10))
}

func syncDataByHandleQueue(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go loopOverQueueAsset(&waitGroup,handleAssetUpdate)
	waitGroup.Wait()


	//test code
	for key,value:=range Records {
		fmt.Println(key,value.asset.Owner.String())
	}

}

func Sync(){
	fillWaitQueue()
	syncDataByHandleQueue()
}
