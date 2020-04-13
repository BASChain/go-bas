package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("Market")
var (

	MarketAddr = "0xa26fDE795d1f15768B588Fb6A9342129AC38C648"
	market     *Contract.BasMarket

	conn = Bas_Ethereum.NewConn()
)


func ResetConnAndService(){
	conn.Reset()
	market = nil
}

func BasMarket() *Contract.BasMarket{
	if market==nil{
		if m,err:=Contract.NewBasMarket(common.HexToAddress(MarketAddr), conn.GetClient());err==nil{
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