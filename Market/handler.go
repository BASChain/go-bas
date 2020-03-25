package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/op/go-logging"
	"sync"
)

var logger, _ = logging.GetLogger("Market")

func loopOverSoldBySell()  {

}

func loopOverSoldByAsk()  {

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
	event:=d.(Contract.BasMarketSellAdded)
	updateSellOrdersByEvent(
		event.Operator,
		event.NameHash,
		*event.Price,
		&event.Raw.BlockNumber)
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
	event:=d.(Contract.BasMarketSellChanged)
	updateSellOrdersByEvent(
		event.Operator,
		event.NameHash,
		*event.Price,
		&event.Raw.BlockNumber)
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
