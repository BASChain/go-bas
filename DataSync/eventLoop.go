package DataSync

import (
	"fmt"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"golang.org/x/net/context"
)


type EventIterator interface {
	Next() bool
}
type EventHandler func(interface{})

func LoopEvent(it EventIterator, handle EventHandler){
	for it.Next() {
		handle(it)
	}
}


func mintAssetHandler(event interface{}){
	it:=event.(*Contract.BasAssetMintAssetIterator)
	insertWaitQueue(string(it.Event.Name))
}



func LoopOverMintAsset(){
	var opts  = bind.FilterOpts{
		Start:0,
		End: nil,
		Context:context.Background(),
	}
	it,err:=Bas_Ethereum.BasAsset().FilterMintAsset(&opts)
	if err==nil{
		LoopEvent(it,mintAssetHandler)
	}

	for  _,s := range waitQueue.ToSlice() {
		fmt.Print(s.(string))
	}
}