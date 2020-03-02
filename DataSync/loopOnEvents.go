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

//////////////////////////////////////////////////////////
//               BasAsset.sol                          //
////////////////////////////////////////////////////////

func handleMintAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetMintAssetIterator)
	insertQueueAsset(Bas_Ethereum.GetHash(string(it.Event.Name)))
}

func handleTakeoverAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetTakeoverAssetIterator)
	insertQueueAsset(it.Event.Hash)
}

func handleRechargeAsset(iterator interface{}){
	it:=iterator.(*Contract.BasAssetRechargeAssetIterator)
	insertQueueAsset(it.Event.Hash)
}

func handleRootChanged(iterator interface{}){
	it:=iterator.(*Contract.BasAssetRootChangedIterator)
	insertQueueAsset(it.Event.NameHash)
}

func handleDNSRecordChange(iterator interface{}){
	it:=iterator.(*Contract.BasAssetDNSRecordChangeIterator)
	insertQueueDns(it.Event.NameHash)
}

func handleDNSRecordRemove(iterator interface{})  {
	it:=iterator.(*Contract.BasAssetDNSRecordRemoveIterator)
	insertQueueDns(it.Event.NameHash)
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