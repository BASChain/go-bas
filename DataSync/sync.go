package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)



func fillWaitQueue(lastBlockNumber uint64){
	opts := getLoopOpts(lastSavingPoint,&lastBlockNumber)
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	go loopOverRootChanged(opts,&waitGroup)
	go loopOverSubChanged(opts,&waitGroup)
	go loopOverDNSChanged(opts,&waitGroup)
	go loopOverAdd(opts,&waitGroup)
	go loopOverUpdate(opts,&waitGroup)
	go loopOverTakeover(opts,&waitGroup)
	go loopOverTransfer(opts,&waitGroup)
	go loopOverTransferFrom(opts,&waitGroup)
	go loopOverRemove(opts,&waitGroup)
	go loopOverPaid(opts,&waitGroup)

	waitGroup.Wait()

}

func syncDataByHandleQueue(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go loopOverQueueOwnership(&waitGroup)
	go loopOverQueueRoot(&waitGroup)
	go loopOverQueueSub(&waitGroup)
	go loopOverQueueDns(&waitGroup)

	waitGroup.Wait()

	clearQuery(queueOwnership)
	clearQuery(queueRoot)
	clearQuery(queueSub)
	clearQuery(queueDns)

	ShowCachedNames()
}

func ShowCachedNames(){
	if DebugFlag{
		for _,v := range Records{
			logger.Info(string(v.Name), v.Owner.String())
		}
	}
}

var subs []event.Subscription

func watch(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup

	opts := getWatchOpts(lastBlockNumber)
	subs = []event.Subscription{}

	waitGroup.Add(10)

	go watchRootChanged(opts,&subs,&waitGroup)
	go watchSubChanged(opts,&subs,&waitGroup)
	go watchDNSChanged(opts,&subs,&waitGroup)
	go watchAdd(opts,&subs,&waitGroup)
	go watchUpdate(opts,&subs,&waitGroup)
	go watchTakeover(opts,&subs,&waitGroup)
	go watchTransfer(opts,&subs,&waitGroup)
	go watchTransferFrom(opts,&subs,&waitGroup)
	go watchRemove(opts,&subs,&waitGroup)
	go watchPaid(opts,&subs,&waitGroup)

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
