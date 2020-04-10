package Market

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



func syncGap(from,to uint64){
	if from>to {
		return
	}
	opts := Bas_Ethereum.GetLoopOpts(from,&to)
	var waitGroup sync.WaitGroup
	waitGroup.Add(8)

	go loopOverAskAdded(opts,&waitGroup)
	go loopOverAskChanged(opts,&waitGroup)
	go loopOverAskRemoved(opts,&waitGroup)
	go loopOverSellAdded(opts,&waitGroup)
	go loopOverSellChanged(opts,&waitGroup)
	go loopOverSellRemoved(opts,&waitGroup)
	go loopOverSoldByAsk(opts,&waitGroup)
	go loopOverSoldBySell(opts,&waitGroup)

	waitGroup.Wait()
	loopOverEventQueue()
}

func watchLogic(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup
	opts := Bas_Ethereum.GetWatchOpts(lastBlockNumber)
	waitGroup.Add(8)
	go watchAskAdded(opts,&waitGroup)
	go watchAskChanged(opts,&waitGroup)
	go watchAskRemove(opts,&waitGroup)
	go watchSellAdded(opts,&waitGroup)
	go watchSellChanged(opts,&waitGroup)
	go watchSellRemove(opts,&waitGroup)
	go watchSoldByAsk(opts,&waitGroup)
	go watchSoldBySell(opts,&waitGroup)

	waitGroup.Wait()
}