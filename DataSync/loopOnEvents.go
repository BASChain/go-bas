package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"sync"
)

var logger, _ = logging.GetLogger("DataSync")

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



func loopOverMintAsset(waitGroup *sync.WaitGroup,opts *bind.FilterOpts, handler EventHandler){
	it,err:=Bas_Ethereum.BasAsset().FilterMintAsset(opts)
	if err==nil{
		LoopEvent(it, handler)
	}
	waitGroup.Done()
}







