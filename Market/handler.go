package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/BASChain/go-bas/Notification"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"sync"
)

var logger, _ = logging.GetLogger("Market")

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

func loopOverSellAdded(opts *bind.FilterOpts,wg *sync.WaitGroup) {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterSellAdded(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"SellAdded",
				it.Event)
		}
	}
}

func handleSellAdded(d interface{}){
	event:=d.(*Contract.BasMarketSellAdded)
	insertSellOrders(
		event.Operator,
		event.NameHash,
		*event.Price,
		&event.Raw.BlockNumber)
	logger.Info("sell add ", echoSellOrder(event.Operator,event.NameHash))

}

func watchSellAdded(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketSellAdded)
	sub,err:=Bas_Ethereum.BasMarket().WatchSellAdded(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sell added")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sell added runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleSellAdded(log)
			}
		}
	}else{
		logger.Error("subscript sell add failed " ,err)
	}
}

func loopOverSellChanged(opts *bind.FilterOpts,wg *sync.WaitGroup)  {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterSellChanged(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"SellChanged",
				it.Event)
		}
	}
}

func handleSellChanged(d interface{}){
	event:=d.(*Contract.BasMarketSellChanged)
	updateSellOrders(
		event.Operator,
		event.NameHash,
		*event.Price)
	logger.Info("sell changed ", echoSellOrder(event.Operator,event.NameHash))
}

func watchSellChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketSellChanged)
	sub,err:=Bas_Ethereum.BasMarket().WatchSellChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sell changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sell changed runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleSellChanged(log)
			}
		}
	}else{
		logger.Error("subscript sell changed failed " ,err)
	}
}

func loopOverSellRemoved(opts *bind.FilterOpts,wg *sync.WaitGroup) {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterSellRemoved(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"SellRemoved",
				it.Event)
		}
	}
}

func handleSellRemoved(d interface{})  {
	event:=d.(*Contract.BasMarketSellRemoved)
	logger.Info("sell remove ", echoSellOrder(event.Operator,event.NameHash))
	removeSellOrders(event.Operator,event.NameHash)
}

func watchSellRemove(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketSellRemoved)
	sub,err:=Bas_Ethereum.BasMarket().WatchSellRemoved(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sell remove")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sell remove runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleSellRemoved(log)
			}
		}
	}else{
		logger.Error("subscript sell remove failed " ,err)
	}
}

func loopOverAskAdded(opts *bind.FilterOpts,wg *sync.WaitGroup) {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterAskAdded(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"AskAdded",
				it.Event)
		}
	}
}

func watchAskAdded(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketAskAdded)
	sub,err:=Bas_Ethereum.BasMarket().WatchAskAdded(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching ask added")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript ask added runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleAskAdded(log)
			}
		}
	}else{
		logger.Error("subscript ask add failed " ,err)
	}
}

func handleAskAdded(d interface{}){
	event:=d.(*Contract.BasMarketAskAdded)
	insertAskOrders(
		event.Operator,
		event.NameHash,
		*event.Price,
		*event.Time,
		&event.Raw.BlockNumber)
	logger.Info("ask add ", echoAskOrder(event.Operator,event.NameHash))
}

func loopOverAskChanged(opts *bind.FilterOpts,wg *sync.WaitGroup) {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterAskChanged(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"AskChanged",
				it.Event)
		}
	}
}

func handleAskChanged(d interface{}){
	event:=d.(*Contract.BasMarketAskChanged)
	updateAskOrders(
		event.Operator,
		event.NameHash,
		*event.Price,
		*event.Time)
	logger.Info("ask changed ", echoAskOrder(event.Operator,event.NameHash))
}

func watchAskChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketAskChanged)
	sub,err:=Bas_Ethereum.BasMarket().WatchAskChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching ask changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript ask changed runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleAskChanged(log)
			}
		}
	}else{
		logger.Error("subscript ask changed failed " ,err)
	}
}

func loopOverAskRemoved(opts *bind.FilterOpts,wg *sync.WaitGroup)  {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterAskRemoved(opts)
	if err==nil{
		for it.Next() {
			insertEq(it.Event.Raw.BlockNumber,
				it.Event.Raw.TxIndex,
				"AskRemoved",
				it.Event)
		}
	}
}

func handleAskRemoved(d interface{})  {
	event:=d.(*Contract.BasMarketAskRemoved)
	logger.Info("ask remove ", echoAskOrder(event.Operator,event.NameHash))
	removeAskOrders(event.Operator,event.NameHash)
}

func watchAskRemove(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketAskRemoved)
	sub,err:=Bas_Ethereum.BasMarket().WatchAskRemoved(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching ask remove")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript ask remove runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				handleAskRemoved(log)
			}
		}
	}else{
		logger.Error("subscript ask remove failed " ,err)
	}
}

func loopOverSoldBySell(opts *bind.FilterOpts,wg *sync.WaitGroup)  {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterSoldBySell(opts)
	if err==nil{
		for it.Next() {
			deal := Deal{it.Event.NameHash,
				BuyFromSell,
				it.Event.From,
				it.Event.To,
				*it.Event.Price,
				it.Event.Raw.BlockNumber}
			Sold = append(Sold, deal)
			logger.Info(echoSold(deal))

			Notification.NotifyMarketSoldBySell(
				it.Event.NameHash,
				it.Event.From,
				it.Event.To,
				*it.Event.Price,
				it.Event.Raw.BlockNumber)
		}
	}
}

func watchSoldBySell(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketSoldBySell)
	sub,err:=Bas_Ethereum.BasMarket().WatchSoldBySell(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sold by sell")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sold by sell runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				deal := Deal{log.NameHash,
					BuyFromSell,
					log.From,
					log.To,
					*log.Price,
					log.Raw.BlockNumber}
				Sold = append(Sold, deal)
				delete(SellOrders[log.From],log.NameHash)
				delete(AskOrders[log.To],log.NameHash)
				logger.Info(echoSold(deal))

				Notification.NotifyMarketSoldBySell(
					log.NameHash,
					log.From,
					log.To,
					*log.Price,
					log.Raw.BlockNumber)
			}
		}
	}else{
		logger.Error("subscript sold by sell failed " ,err)
	}
}

func loopOverSoldByAsk(opts *bind.FilterOpts,wg *sync.WaitGroup)  {
	defer wg.Done()
	it,err:=Bas_Ethereum.BasMarket().FilterSoldByAsk(opts)
	if err==nil{
		for it.Next() {
			deal := Deal{it.Event.NameHash,
				SellToAsk,
				it.Event.From,
				it.Event.To,
				*it.Event.Price,
				it.Event.Raw.BlockNumber}
			Sold = append(Sold, deal)
			delete(SellOrders[it.Event.From],it.Event.NameHash)
			delete(AskOrders[it.Event.To],it.Event.NameHash)
			logger.Info(echoSold(deal))

			Notification.NotifyMarketSoldByAsk(
				it.Event.NameHash,
				it.Event.From,
				it.Event.To,
				*it.Event.Price,
				it.Event.Raw.BlockNumber)
		}
	}
}

func watchSoldByAsk(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasMarketSoldByAsk)
	sub,err:=Bas_Ethereum.BasMarket().WatchSoldByAsk(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching sold by ask")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript sold by ask runtime error", e)
				return
			case log:= <-logs:
				lastSavingPoint = log.Raw.BlockNumber
				deal := Deal{log.NameHash,
					SellToAsk,
					log.From,
					log.To,
					*log.Price,
					log.Raw.BlockNumber}
				Sold = append(Sold, deal)
				delete(SellOrders[log.From],log.NameHash)
				delete(AskOrders[log.To],log.NameHash)
				logger.Info(echoSold(deal))

				Notification.NotifyMarketSoldByAsk(
					log.NameHash,
					log.From,
					log.To,
					*log.Price,
					log.Raw.BlockNumber)
			}
		}
	}else{
		logger.Error("subscript sold by ask failed " ,err)
	}
}
