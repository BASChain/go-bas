package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

func fillEventQueue(lastBlockNumber uint64){
	opts := getLoopOpts(lastSavingPoint,&lastBlockNumber)
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
}

func syncDataByHandleQueue(){
	loopOverEventQueue()
}

var subs []event.Subscription


func watch(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup

	opts := getWatchOpts(lastBlockNumber)
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
	Bas_Ethereum.ResetConnAndContracts()
	Sync()
}

func Sync(){
	lastBlockNumber := Bas_Ethereum.GetLastBlockNumber()
	fillEventQueue(lastBlockNumber)
	syncDataByHandleQueue()
	lastSavingPoint = lastBlockNumber
	logger.Info("current saving point is blocknumber : ", lastSavingPoint)
	watch(lastBlockNumber)
}