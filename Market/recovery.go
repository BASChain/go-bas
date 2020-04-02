package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_market = "0xA32ccce4B7aB28d3Ce40BBa03A2748bCbe4544dB"
	market *Contract.BasMarket

	conn = Bas_Ethereum.NewConn()
)


func ResetConnAndService(){
	conn.Reset()
	market = nil
}

func BasMarket() *Contract.BasMarket{
	if market==nil{
		if m,err:=Contract.NewBasMarket(common.HexToAddress(_market), conn.GetClient());err==nil{
			market = m
		}else{
			logger.Fatal("can't recover market",err)
		}
	}
	return market
}

func BlockNumnber2TimeStamp(blockNumber uint64) int64  {
	t,_:=conn.GetTimestamp(blockNumber)

	return int64(t)
}