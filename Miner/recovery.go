package Miner


import (
"github.com/BASChain/go-bas/Bas_Ethereum"
Contract "github.com/BASChain/go-bas/Contracts"
"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("Miner")

var (
	MinerAddr = "0xb685C02bF992c61c68393aF7fcD8F46833Fb6937"

	miner *Contract.BasMiner

	conn = Bas_Ethereum.NewConn()
)


func ResetConnAndService(){
	conn.Reset()
	miner = nil
}

func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(MinerAddr), conn.GetClient());err==nil{
			miner = m
		}else{
			logger.Fatal("can't recover miner",err)
		}
	}
	return miner
}