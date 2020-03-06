package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"sync"
)

var logger, _ = logging.GetLogger("DataSync")

type EventHandler func(iterator interface{})
type updateHandler func(key interface{})

type EventIterator interface {
	Next() bool
}

func LoopEvent(it EventIterator, handle EventHandler){
	for it.Next() {
		handle(it)
	}
}

func getLoopOpts(s uint64, e *uint64)  *bind.FilterOpts{
	var opts  = &bind.FilterOpts{
		Start:s,
		End: e,
		Context:context.Background(),
	}
	return opts
}

func getWatchOpts(s uint64) *bind.WatchOpts{
	var opts = &bind.WatchOpts{
		Start:   &s,
		Context: context.Background(),
	}
	return opts
}

//////////////////////////////////////////////////////////
//               BasAsset.sol                          //
////////////////////////////////////////////////////////

func handleMintAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetMintAssetIterator)
	insertQueueAsset(Bas_Ethereum.GetHash(string(it.Event.Name)))
}

func watchMintAsset(opts *bind.WatchOpts){
	logs := make(chan *Contract.BasAssetMintAsset)
	sub,err:=Bas_Ethereum.BasAsset().WatchMintAsset(opts,logs)
	if err==nil{
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript MintAsset runtime error", e)
				return
			case log:= <-logs:
				handleAssetUpdate(Bas_Ethereum.GetHash(string(log.Name)))
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript MintAsset failed " ,err)
	}
}

func handleTakeoverAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetTakeoverAssetIterator)
	insertQueueAsset(it.Event.Hash)
}

func watchTakeoverAsset(opts *bind.WatchOpts){
	logs:= make(chan *Contract.BasAssetTakeoverAsset)
	sub,err := Bas_Ethereum.BasAsset().WatchTakeoverAsset(opts,logs)
	if err==nil{
		for {
			select {
			case err:=<-sub.Err():
				logger.Error("subscript Takeover runtime error",err)
				return
			case log:=<-logs:
				handleAssetUpdate(log.Hash)
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript Takeover failed " ,err)
	}
}

func handleRechargeAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetRechargeAssetIterator)
	insertQueueAsset(it.Event.Hash)
}

func watchRechargeAsset(opts *bind.WatchOpts){
	logs := make(chan *Contract.BasAssetRechargeAsset)
	sub,err:=Bas_Ethereum.BasAsset().WatchRechargeAsset(opts,logs)
	if err==nil{
		for {
			select {
			case err:=<-sub.Err():
				logger.Error("subscript Recharge asset runtime error",err)
				return
			case log:= <-logs:
				handleAssetUpdate(log.Hash)
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript Recharge asset failed " ,err)
	}
}

func handleRootChanged(iterator interface{}){
	it:=iterator.(*Contract.BasAssetRootChangedIterator)
	insertQueueAsset(it.Event.NameHash)
}

func watchRootChanged(opts *bind.WatchOpts){
	logs := make(chan *Contract.BasAssetRootChanged)
	sub,err:=Bas_Ethereum.BasAsset().WatchRootChanged(opts,logs)
	if err==nil{
		for {
			select {
			case err:=<-sub.Err():
				logger.Error("subscript root changed runtime error",err)
				return
			case log:= <-logs:
				handleAssetUpdate(log.NameHash)
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript watch root changed failed " ,err)
	}
}


func handleDNSRecordChange(iterator interface{}){
	it:=iterator.(*Contract.BasAssetDNSRecordChangeIterator)
	insertQueueDns(it.Event.NameHash)
}

func watchDNSRecordChange(opts *bind.WatchOpts){
	logs :=make(chan *Contract.BasAssetDNSRecordChange)
	sub,err:=Bas_Ethereum.BasAsset().WatchDNSRecordChange(opts,logs)
	if err==nil{
		for {
			select {
			case err:=<-sub.Err():
				logger.Error("subscript DNSRecord changed runtime error",err)
				return
			case log:= <-logs:
				handleDNSUpdate(log.NameHash)
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript DNSRecord change failed " ,err)
	}
}

func handleDNSRecordRemove(iterator interface{})  {
	it:=iterator.(*Contract.BasAssetDNSRecordRemoveIterator)
	insertQueueDns(it.Event.NameHash)
}

func watchDNSRecordRemove(opts *bind.WatchOpts){
	logs :=make(chan *Contract.BasAssetDNSRecordRemove)
	sub,err:=Bas_Ethereum.BasAsset().WatchDNSRecordRemove(opts,logs)
	if err==nil{
		for {
			select {
			case err:=<-sub.Err():
				logger.Error("subscript DNSRecord remove runtime error",err)
				return
			case log:= <-logs:
				handleDNSUpdate(log.NameHash)
				lastSavingPoint = log.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("subscript DNSRecord remove failed " ,err)
	}
}


func handleAssetUpdate(key interface{}){
	updateAsset(key.(Bas_Ethereum.Hash))
}

func handleDNSUpdate(key interface{}){
	updateDNS(key.(Bas_Ethereum.Hash))
}


func loopOverMintAsset(opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterMintAsset(opts)
	if err==nil{
		LoopEvent(it, handler)
	}else{
		logger.Error("loop over mint asset err :" , err)
	}
}


func loopOverTakeoverAsset(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterTakeoverAsset(opts)
	if err==nil{
		LoopEvent(it,handler)
	}else{
		logger.Error("loop over takeover asset err :" , err)
	}
	waitGroup.Done()
}

func loopOverRechargeAsset(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterRechargeAsset(opts)
	if err==nil{
		LoopEvent(it, handler)
	}else{
		logger.Error("loop over recharge asset err :" , err)
	}
	waitGroup.Done()
}

func loopOverRootChanged(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterRootChanged(opts)
	if err==nil{
		LoopEvent(it,handler)
	}else{
		logger.Error("loop over root changed err :" , err)
	}
	waitGroup.Done()
}

func loopOverDNSRecordChange(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterDNSRecordChange(opts)
	if err==nil{
		LoopEvent(it,handler)
	}else{
		logger.Error("loop over DNS Record change err :" , err)
	}
	waitGroup.Done()
}

func loopOverDNSRecordRemove(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterDNSRecordRemove(opts)
	if err==nil{
		LoopEvent(it,handler)
	}else{
		logger.Error("loop over DNS Record remove err :" , err)
	}
	waitGroup.Done()
}