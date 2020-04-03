package Miner

import (
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"sync"
)

func loopOverPaid(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasMiner().FilterAllocationChanged(opts)
	if err==nil{
		for it.Next() {





		}
	}else{
		logger.Error("loop over paid err :" , err)
	}
}

func handleAllocation(alc *Contract.BasMinerAllocationChanged)  {
	switch alc.AllocateType {
	case 0:
	case 1:
	case 2:
	case 3:
	default:
	}
}



//func watchPaid(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
//	logs := make(chan *Contract.BasOANNPaid)
//	sub,err:=BasOANN().WatchPaid(opts,logs)
//	defer wg.Done()
//	if err==nil{
//		logger.Info("watching paid")
//		*subs = append(*subs, sub)
//		for {
//			select {
//			case e :=<-sub.Err():
//				logger.Error("subscript paid runtime error", e)
//				return
//			case log:= <-logs:
//				SyncGapWithNoTrust(log.Raw.BlockNumber)
//				updatePaid(Receipt{
//					Payer:  log.Payer,
//					Name:   log.Name,
//					Option: log.Option,
//					Amount: log.Amount,
//					CommitBlock:log.Raw.BlockNumber,
//				},log.Receipt)
//				logger.Info("detected address ",
//					log.Payer.String(), " paid " , log.Amount.String(), "for ", log.Option," on", string(log.Name))
//			}
//		}
//	}else{
//		logger.Error("subscript paid failed " ,err)
//	}
//}
