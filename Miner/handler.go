package Miner

import (
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

const (
	AllocationRoot uint8 = 0
	AllocationSub uint8 = 1
	AllocationSelfSub uint8 = 2
	AllocationCustomedSub uint8 = 3
)


func loopOverAllocationChanged(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterAllocationChanged(opts)
	if err==nil{
		for it.Next() {
			eq.Insert(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"AllocationChanged",
				it.Event)
		}
	}else{
		logger.Error("loop over paid err :" , err)
	}
}

func getRecord(blockNumber uint64, index uint) Setting {
	var rcd Setting
	err,d:=SettingRecords.GetClosest(blockNumber, index)
	if err !=nil {
		logger.Info(err)
		rcd = Setting{}
	}else{
		rcd = d.(Setting)
	}
	return  rcd
}

func handleAllocation(d interface{})  {
	alc := d.(*Contract.BasMinerAllocationChanged)
	blockNumber := alc.Raw.BlockNumber
	index := alc.Raw.TxIndex
	rcd := getRecord(blockNumber,index)
	switch alc.AllocateType {
	case AllocationRoot:
		rcd.Allocation[0] = *alc.ToAdmin
		rcd.Allocation[1] = *alc.ToBurn
		rcd.Allocation[2] = *alc.ToMiner
	case AllocationSub:
		rcd.Allocation[3] = *alc.ToAdmin
		rcd.Allocation[4] = *alc.ToBurn
		rcd.Allocation[5] = *alc.ToMiner
		rcd.Allocation[6] = *alc.ToRoot
	case AllocationSelfSub:
		rcd.Allocation[7] = *alc.ToAdmin
		rcd.Allocation[8] = *alc.ToBurn
		rcd.Allocation[9] = *alc.ToMiner
		rcd.Allocation[10] = *alc.ToRoot
	case AllocationCustomedSub:
		rcd.Allocation[11] = *alc.ToAdmin
		rcd.Allocation[12] = *alc.ToBurn
		rcd.Allocation[13] = *alc.ToMiner
		rcd.Allocation[14] = *alc.ToRoot
	default:
		logger.Error("other allocate type :", alc.AllocateType)
		return
	}
	logger.Info("[IMPORTANT] handle allocation change, type : ", alc.AllocateType,
		" {", (*alc.ToAdmin).String(), (*alc.ToBurn).String(), (*alc.ToMiner).String(), (*alc.ToRoot).String(), "}" )
	SettingRecords.Insert(blockNumber,index,rcd)
}



func watchAllocationChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerAllocationChanged)
	sub,err:=BasMiner().WatchAllocationChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching allocation changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript allocation changed runtime error", e)
				return
			case log:= <-logs:
				handleAllocation(log)
			}
		}
	}else{
		logger.Error("subscript allocation changed failed " ,err)
	}
}

func loopOverMinerAdd(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterMinerAdd(opts)
	if err==nil{
		for it.Next() {
			eq.Insert(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"MinerAdd",
				it.Event)
		}
	}else{
		logger.Error("loop over miner add err :" , err)
	}
}

func handleMinerAdd(d interface{}) {
	e := d.(*Contract.BasMinerMinerAdd)
	blockNumber := e.Raw.BlockNumber
	index := e.Raw.TxIndex
	rcd := getRecord(blockNumber,index)
	rcd.Miners = append(rcd.Miners, e.Miner)
	SettingRecords.Insert(blockNumber,index,rcd)
	logger.Info("[IMPORTANT] miner added : ", e.Miner.String())
}

func watchMinerAdd(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerMinerAdd)
	sub,err:=BasMiner().WatchMinerAdd(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching miner add")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript miner add runtime error", e)
				return
			case log:= <-logs:
				handleMinerAdd(log)
			}
		}
	}else{
		logger.Error("subscript miner add failed " ,err)
	}
}

func loopOverMinerRemove(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterMinerRemove(opts)
	if err==nil{
		for it.Next() {
			eq.Insert(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"MinerRemove",
				it.Event)
		}
	}else{
		logger.Error("loop over miner remove err :" , err)
	}
}

func handleMinerRemove(d interface{}) {
	e := d.(*Contract.BasMinerMinerRemove)
	blockNumber := e.Raw.BlockNumber
	index := e.Raw.TxIndex
	rcd := getRecord(blockNumber,index)

	for i:= 0; i < len(rcd.Miners) ; i++ {
		if rcd.Miners[i] == e.Miner {
			rcd.Miners = append(rcd.Miners[:i],rcd.Miners[i+1:]...)
			logger.Info("[IMPORTANT] miner removed : ", e.Miner.String())
			SettingRecords.Insert(blockNumber,index,rcd)
			return
		}
	}
	logger.Error("[IMPORTANT] miner removed not found! miner : ", e.Miner.String())
}

func watchMinerRemove(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerMinerRemove)
	sub,err:=BasMiner().WatchMinerRemove(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching miner remove")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript miner remove runtime error", e)
				return
			case log:= <-logs:
				handleMinerRemove(log)
			}
		}
	}else{
		logger.Error("subscript miner remove failed " ,err)
	}
}

func loopOverMinerReplace(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterMinerReplace(opts)
	if err==nil{
		for it.Next() {
			eq.Insert(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"MinerReplace",
				it.Event)
		}
	}else{
		logger.Error("loop over miner replace err :" , err)
	}
}

func handleMinerReplace(d interface{}) {
	e := d.(*Contract.BasMinerMinerReplace)
	blockNumber := e.Raw.BlockNumber
	index := e.Raw.TxIndex
	rcd := getRecord(blockNumber,index)

	for i:= 0; i < len(rcd.Miners) ; i++ {
		if rcd.Miners[i] == e.OldMiner {
			rcd.Miners[i] = e.NewMiner
			logger.Info("[IMPORTANT] miner replaced from :", e.OldMiner.String(), "to :",e.NewMiner.String())
			SettingRecords.Insert(blockNumber,index,rcd)
			return
		}
	}
	logger.Error("[IMPORTANT] miner replace not found! miner : ", e.OldMiner.String())
}

func watchMinerReplace(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerMinerReplace)
	sub,err:=BasMiner().WatchMinerReplace(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching miner replace")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript miner replace runtime error", e)
				return
			case log:= <-logs:
				handleMinerReplace(log)
			}
		}
	}else{
		logger.Error("subscript miner replace failed " ,err)
	}
}

func loopOverReceipt(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterReceipt(opts)
	if err==nil{
		for it.Next() {
			handleReceipt(it.Event)
		}
	}else{
		logger.Error("loop over receipt err :" , err)
	}
}

func handleReceipt(r *Contract.BasMinerReceipt){
	Receipts = append(Receipts, SimplifiedRecepit{
		BlockNumber:   r.Raw.BlockNumber,
		TxIndex:       r.Raw.TxIndex,
		ReceiptNumber: r.ReceiptNumber,
		Amount:        *r.Amout,
		Allocation:    r.Allocation,
	})
}

func watchReceipt(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerReceipt)
	sub,err:=BasMiner().WatchReceipt(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching miner receipt")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript miner receipt runtime error", e)
				return
			case log:= <-logs:
				handleReceipt(log)
			}
		}
	}else{
		logger.Error("subscript miner receipt failed " ,err)
	}
}

func loopOverWithdraw(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterWithdraw(opts)
	if err==nil{
		for it.Next() {
			handleWithdraw(it.Event)
		}
	}else{
		logger.Error("loop over withdraw err :" , err)
	}
}

func handleWithdraw(w *Contract.BasMinerWithdraw){
	Withdraws = append(Withdraws, SimplifiedWithdraw{
		Drawer: w.Drawer,
		Amount: *w.Amout,
	})
}

func watchWithdraw(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMinerWithdraw)
	sub,err:=BasMiner().WatchWithdraw(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching miner withdraw")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript miner withdraw runtime error", e)
				return
			case log:= <-logs:
				handleWithdraw(log)
			}
		}
	}else{
		logger.Error("subscript miner withdraw failed " ,err)
	}
}