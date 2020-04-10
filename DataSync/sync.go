package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sync"
	"time"
)

var Sync *Bas_Ethereum.SyncHelper = Bas_Ethereum.NewSyncHelper(
	conn,
	syncGap,
	watchLogic,
	ResetConnAndService,
	func() {
		Settings()
		go func() {
			time.Sleep(time.Duration(30) * time.Second)
			ShowCachedNames()
		}()

	},
)


func syncGap(from ,to uint64){
	if from>to {
		return
	}
	opts := Bas_Ethereum.GetLoopOpts(from,&to)
	var waitGroup sync.WaitGroup
	waitGroup.Add(12)

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
	go loopOverSettingChanged(opts,&waitGroup)

	waitGroup.Wait()
	syncDataByHandleQueue()
}


func syncDataByHandleQueue(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go loopOverQueueOwnership(&waitGroup)
	go loopOverQueueRoot(&waitGroup)
	go loopOverQueueSub(&waitGroup)
	go loopOverQueueDns(&waitGroup)

	waitGroup.Wait()

	clearQueueO()
	clearQueueD()
	clearQueueR()
	clearQueueS()
}

func ShowCachedNames(){
	if DebugFlag{
		for _,v := range Records{
			logger.Info(string(v.Name), v.Owner.String())
		}
	}
}

func watchLogic(lastBlockNumber uint64){
	var waitGroup sync.WaitGroup
	opts := Bas_Ethereum.GetWatchOpts(lastBlockNumber)
	waitGroup.Add(12)
	go watchRootChanged(opts,&waitGroup)
	go watchSubChanged(opts,&waitGroup)
	go watchDNSChanged(opts,&waitGroup)
	go watchAdd(opts,&waitGroup)
	go watchUpdate(opts,&waitGroup)
	go watchExtend(opts,&waitGroup)
	go watchTakeover(opts,&waitGroup)
	go watchTransfer(opts,&waitGroup)
	go watchTransferFrom(opts,&waitGroup)
	go watchRemove(opts,&waitGroup)
	go watchPaid(opts,&waitGroup)
	go watchSettingChanged(opts,&waitGroup)

	waitGroup.Wait()
}
