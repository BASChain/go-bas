package Miner


import (
"github.com/BASChain/go-bas/Bas_Ethereum"
Contract "github.com/BASChain/go-bas/Contracts"
"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("Miner")

var (
	_miner = "0xCAB59645aE535A7b5a4f81d8D17E2fe0d2Cf4687"

	miner *Contract.BasMiner

	conn = Bas_Ethereum.NewConn()
)


func ResetConnAndService(){
	conn.Reset()
	miner = nil
}

func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(_miner), conn.GetClient());err==nil{
			miner = m
		}else{
			logger.Fatal("can't recover miner",err)
		}
	}
	return miner
}