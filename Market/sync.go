package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

var lastSavingPoint = uint64(0)
var currentSavingPoint = uint64(0)

func moveToNewSavingPoint(blockNumber uint64) bool{
	if currentSavingPoint == blockNumber {
		return false
	}
	lastSavingPoint = currentSavingPoint;
	currentSavingPoint = blockNumber;
	logger.Info("[Market]  saving point  ", lastSavingPoint, "----->" , currentSavingPoint)
	return true
}

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

func SyncGapWithNoTrust(blockNumber uint64){
	if moveToNewSavingPoint(blockNumber){
		go syncGap(lastSavingPoint+1, blockNumber-1)
	}
}

func syncGapToNewest(){
	moveToNewSavingPoint(conn.GetLastBlockNumber())
	syncGap(lastSavingPoint,currentSavingPoint)
}

var subs []event.Subscription


func watch(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup

	opts := Bas_Ethereum.GetWatchOpts(lastBlockNumber)
	subs = []event.Subscription{}

	waitGroup.Add(8)

	go watchAskAdded(opts,&subs,&waitGroup)
	go watchAskChanged(opts,&subs,&waitGroup)
	go watchAskRemove(opts,&subs,&waitGroup)
	go watchSellAdded(opts,&subs,&waitGroup)
	go watchSellChanged(opts,&subs,&waitGroup)
	go watchSellRemove(opts,&subs,&waitGroup)
	go watchSoldByAsk(opts,&subs,&waitGroup)
	go watchSoldBySell(opts,&subs,&waitGroup)

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
	ResetConnAndService()
	watch(currentSavingPoint)
}

func Sync(){
	syncGapToNewest()
	watch(currentSavingPoint)
}