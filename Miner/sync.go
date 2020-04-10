package Miner

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
)



var Sync *Bas_Ethereum.SyncHelper = Bas_Ethereum.NewSyncHelper(
	conn,
	syncGap,
	watchLogic,
	ResetConnAndService,
	func() {},
)

func syncGap(from ,to uint64){
	if from>to {
		return
	}
	opts := Bas_Ethereum.GetLoopOpts(from,&to)
	var waitGroup sync.WaitGroup
	waitGroup.Add(6)

	go loopOverAllocationChanged(opts,&waitGroup)
	go loopOverMinerAdd(opts,&waitGroup)
	go loopOverMinerRemove(opts,&waitGroup)
	go loopOverMinerReplace(opts,&waitGroup)
	go loopOverReceipt(opts,&waitGroup)
	go loopOverWithdraw(opts,&waitGroup)

	waitGroup.Wait()
	loopOverEventQueue()
}

func watchLogic(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup
	opts := Bas_Ethereum.GetWatchOpts(lastBlockNumber)
	waitGroup.Add(6)
	go watchAllocationChanged(opts,&waitGroup)
	go watchMinerAdd(opts,&waitGroup)
	go watchMinerRemove(opts,&waitGroup)
	go watchMinerReplace(opts,&waitGroup)
	go watchReceipt(opts,&waitGroup)
	go watchWithdraw(opts,&waitGroup)

	waitGroup.Wait()
}
