package Miner

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

var lastSavingPoint = uint64(0)
var currentSavingPoint = uint64(0)

func moveToNewSavingPoint(blockNumber uint64){
	if currentSavingPoint == blockNumber {
		return
	}
	lastSavingPoint = currentSavingPoint;
	currentSavingPoint = blockNumber;
	logger.Info("[Miner]  saving point  ", lastSavingPoint, "----->" , currentSavingPoint)
}

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

func SyncGapWithNoTrust(blockNumber uint64){
	moveToNewSavingPoint(blockNumber)
	syncGap(lastSavingPoint+1,blockNumber-1)
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

	waitGroup.Add(6)

	go watchAllocationChanged(opts,&subs,&waitGroup)
	go watchMinerAdd(opts,&subs,&waitGroup)
	go watchMinerRemove(opts,&subs,&waitGroup)
	go watchMinerReplace(opts,&subs,&waitGroup)
	go watchReceipt(opts,&subs,&waitGroup)
	go watchWithdraw(opts,&subs,&waitGroup)

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

func ShowSysParams(){
	_,rcd:=SettingRecords.GetClosest(currentSavingPoint,0)
	st:=rcd.(Setting)
	logger.Info("allocations root: ", st.Allocation)
	for _,m := range st.Miners {
		logger.Info("miner -->" ,m.String())
	}
}

var firstStart = true;

func Sync(){
	syncGapToNewest()
	if firstStart {
		ShowSysParams()
		firstStart = false
	}
	watch(currentSavingPoint)
}
