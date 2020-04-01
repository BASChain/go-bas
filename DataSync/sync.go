package DataSync

import (
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

var lastSavingPoint = uint64(0)
var currentSavingPoint = uint64(0)

func moveToNewSavingPoint(blockNumber uint64){
	lastSavingPoint = currentSavingPoint;
	currentSavingPoint = blockNumber;
	logger.Info("[Data_Sync]  saving point  ", lastSavingPoint, "----->" , currentSavingPoint)
}

func syncGap(from ,to uint64){
	if from>to {
		return
	}
	opts := getLoopOpts(from,&to)
	var waitGroup sync.WaitGroup
	waitGroup.Add(11)

	go loopOverRootChanged(opts,&waitGroup)
	go loopOverSubChanged(opts,&waitGroup)
	go loopOverDNSChanged(opts,&waitGroup)
	go loopOverAdd(opts,&waitGroup)
	go loopOverUpdate(opts,&waitGroup)
	go loopOverExtend(opts,&waitGroup)
	go loopOverTakeover(opts,&waitGroup)
	go loopOverTransfer(opts,&waitGroup)
	go loopOverTransferFrom(opts,&waitGroup)
	go loopOverRemove(opts,&waitGroup)
	go loopOverPaid(opts,&waitGroup)

	waitGroup.Wait()
	syncDataByHandleQueue()
}

func SyncGapWithNoTrust(blockNumber uint64){
	moveToNewSavingPoint(blockNumber)
	syncGap(lastSavingPoint+1,blockNumber-1)
}

func syncGapToNewest(){
	moveToNewSavingPoint(conn.GetLastBlockNumber())
	syncGap(lastSavingPoint,currentSavingPoint)
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

	waitGroup.Add(11)

	go watchRootChanged(opts,&subs,&waitGroup)
	go watchSubChanged(opts,&subs,&waitGroup)
	go watchDNSChanged(opts,&subs,&waitGroup)
	go watchAdd(opts,&subs,&waitGroup)
	go watchUpdate(opts,&subs,&waitGroup)
	go watchExtend(opts,&subs,&waitGroup)
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
	ResetConnAndService()
	Sync()
}

var firstStart = true;

func Sync(){
	Settings()
	syncGapToNewest()
	if firstStart {
		ShowCachedNames()
		firstStart = false
	}
	watch(currentSavingPoint)
}
