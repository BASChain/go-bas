package Miner


import (
"github.com/BASChain/go-bas/Bas_Ethereum"
Contract "github.com/BASChain/go-bas/Contracts"
"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("Miner")

var (
	_miner = "0x34149F28b03B2Db2778Cc4F3d99e0ca71C34cc95"
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

func BlockNumnber2TimeStamp(blockNumber uint64) int64  {
	t,_:=conn.GetTimestamp(blockNumber)

	return int64(t)
}