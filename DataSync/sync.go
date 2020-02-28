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

	go loopOverMintAsset(&waitGroup,opts)


	waitGroup.Wait()

	lastSavingPoint = lastBlockNumber
	logger.Info("Collected waitQueue to block number: " + strconv.FormatUint(lastSavingPoint, 10))
}

func syncDataByHandleQueue(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go loopOverQueueAsset(&waitGroup)
	waitGroup.Wait()

	for key,value:=range Records {
		fmt.Print("domian : " +key + ":\n")
		fmt.Println(string(value.asset.Name),value.asset.Owner.String())
	}

}

func Sync(){
	fillWaitQueue()
	syncDataByHandleQueue()
}
