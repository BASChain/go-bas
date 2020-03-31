package DataSync

import (
	"encoding/hex"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"sync"
)

var logger, _ = logging.GetLogger("DataSync")

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

func loopOverRootChanged(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasAsset().FilterRootChanged(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueRoot)
		}
	}else{
		logger.Error("loop over root changed err :" , err)
	}
}

func watchRootChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasAssetRootChanged)
	sub,err:=Bas_Ethereum.BasAsset().WatchRootChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching root changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript root changed runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryRoot(log.NameHash,currentSavingPoint)
				logger.Info("detected root changed : ",
					string(Records[log.NameHash].Name),
				)
			}
		}
	}else{
		logger.Error("subscript root changed failed " ,err)
	}
}

func loopOverSubChanged(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasAsset().FilterSubChanged(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueSub)
		}
	}else{
		logger.Error("loop over sub changed err :" , err)
	}
}

func watchSubChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasAssetSubChanged)
	sub,err:=Bas_Ethereum.BasAsset().WatchSubChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sub changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sub changed runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQuerySub(log.NameHash,currentSavingPoint)
				logger.Info("detected sub changed : ",
					string(Records[log.NameHash].Name),
				)
			}
		}
	}else{
		logger.Error("subscript root changed failed " ,err)
	}
}

func loopOverDNSChanged(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasDNS().FilterDNSChanged(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueDns)
		}
	}else{
		logger.Error("loop over dns changed err :" , err)
	}
}

func watchDNSChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasDNSDNSChanged)
	sub,err:=Bas_Ethereum.BasDNS().WatchDNSChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching dns changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript dns changed runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryDNS(log.NameHash,currentSavingPoint)
				logger.Info("detected dns changed : ",
					string(Records[log.NameHash].Name),
				)
			}
		}
	}else{
		logger.Error("subscript root changed failed " ,err)
	}
}

func loopOverAdd(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterAdd(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
			//logger.Info("!!!add:" , "0x"+hex.EncodeToString(it.Event.NameHash[:]), it.Event.Owner.String())
			if firstAppearInBlock[it.Event.NameHash] == 0 || firstAppearInBlock[it.Event.NameHash] > it.Event.Raw.BlockNumber{
				firstAppearInBlock[it.Event.NameHash] = it.Event.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("loop over add err :" , err)
	}
}

func watchAdd(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipAdd)
	sub,err:=Bas_Ethereum.BasOwnership().WatchAdd(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching add")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript add runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected add : ",
					"0x"+hex.EncodeToString(log.NameHash[:]), "owner : ", log.Owner.String())
			}
		}
	}else{
		logger.Error("subscript add failed " ,err)
	}
}

func loopOverUpdate(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterUpdate(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
			if firstAppearInBlock[it.Event.NameHash] == 0 || firstAppearInBlock[it.Event.NameHash] > it.Event.Raw.BlockNumber{
				firstAppearInBlock[it.Event.NameHash] = it.Event.Raw.BlockNumber
			}
		}
	}else{
		logger.Error("loop over update err :" , err)
	}
}

func watchUpdate(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipUpdate)
	sub,err:=Bas_Ethereum.BasOwnership().WatchUpdate(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching update")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript update runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected update : ",
					"0x"+hex.EncodeToString(log.NameHash[:]), "owner : ", log.Owner.String())
			}
		}
	}else{
		logger.Error("subscript update failed " ,err)
	}
}

func loopOverExtend(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterExtend(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
		}
	}else{
		logger.Error("loop over extend err :" , err)
	}
}

func watchExtend(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipExtend)
	sub,err:=Bas_Ethereum.BasOwnership().WatchExtend(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching extend")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript extend runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected extend : ",
					"0x"+hex.EncodeToString(log.NameHash[:]))
			}
		}
	}else{
		logger.Error("subscript update failed " ,err)
	}
}


func loopOverTakeover(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterTakeover(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
		}
	}else{
		logger.Error("loop over takeover err :" , err)
	}
}

func watchTakeover(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipTakeover)
	sub,err:=Bas_Ethereum.BasOwnership().WatchTakeover(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching takeover")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript takeover runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected takeover : ",
					"0x"+hex.EncodeToString(log.NameHash[:]), "from : ", log.From.String(), "to : ",log.To.String())
			}
		}
	}else{
		logger.Error("subscript takeover failed " ,err)
	}
}

func loopOverTransfer(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterTransfer(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
		}
	}else{
		logger.Error("loop over transfer err :" , err)
	}
}

func watchTransfer(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipTransfer)
	sub,err:=Bas_Ethereum.BasOwnership().WatchTransfer(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching transfer")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript transfer runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected transfer : ",
					"0x"+hex.EncodeToString(log.NameHash[:]), "from : ", log.From.String(), "to : ",log.To.String())
			}
		}
	}else{
		logger.Error("subscript transfer failed " ,err)
	}
}

func loopOverTransferFrom(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterTransferFrom(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
		}
	}else{
		logger.Error("loop over transferFrom err :" , err)
	}
}

func watchTransferFrom(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipTransferFrom)
	sub,err:=Bas_Ethereum.BasOwnership().WatchTransferFrom(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching transferFrom")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript transferFrom runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected transferFrom : ",
					"0x"+hex.EncodeToString(log.NameHash[:]), "from : ", log.From.String(), "to : ",log.To.String())
			}
		}
	}else{
		logger.Error("subscript transferFrom failed " ,err)
	}
}

func loopOverRemove(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOwnership().FilterRemove(opts)
	if err==nil{
		for it.Next() {
			insertQueue(it.Event.NameHash,queueOwnership)
		}
	}else{
		logger.Error("loop over remove err :" , err)
	}
}

func watchRemove(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOwnershipRemove)
	sub,err:=Bas_Ethereum.BasOwnership().WatchRemove(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching remove")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript remove runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updateByQueryOwnership(log.NameHash,currentSavingPoint)
				logger.Info("detected remove : ",
					"0x"+hex.EncodeToString(log.NameHash[:]))
			}
		}
	}else{
		logger.Error("subscript remove failed " ,err)
	}
}

func loopOverPaid(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=Bas_Ethereum.BasOANN().FilterPaid(opts)
	if err==nil{
		for it.Next() {
			updatePaid(Receipt{
				Payer:  it.Event.Payer,
				Name:   it.Event.Name,
				Option: it.Event.Option,
				Amount: it.Event.Amount,
				CommitBlock: it.Event.Raw.BlockNumber,
			})
		}
	}else{
		logger.Error("loop over remove err :" , err)
	}
}

func watchPaid(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOANNPaid)
	sub,err:=Bas_Ethereum.BasOANN().WatchPaid(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching paid")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript paid runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				updatePaid(Receipt{
					Payer:  log.Payer,
					Name:   log.Name,
					Option: log.Option,
					Amount: log.Amount,
					CommitBlock:log.Raw.BlockNumber,
				})
				logger.Info("detected address ",
					log.Payer.String(), " paied " , log.Amount.String(), "for ", log.Option," on", string(log.Name))
			}
		}
	}else{
		logger.Error("subscript paid failed " ,err)
	}
}