package Bas_Ethereum

import (
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)
var logger, _ = logging.GetLogger("Bas_Ethereum")

var (

	AccessPoints = []string{
		"wss://ropsten.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
		"wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
		"ws://75.135.96.248:3334",
	}


	BASTokenAddress    = "0x9d0314f9Bacd569DCB22276867AAEeE1C8A87614"
	BASMinerAddress    = "0x4074594618FFe52Ca9ab03cc01314b1f7893da9D"
	BASAssetAddress    = "0x5346aDb387D87009C133c4773deD55Bbc47A595B"
	BASOANNAddress     = "0x0282C762a66D3BFfbD2E2c6bEe6C56eAfA847453"
	SendFreeBASAddress = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"

	conn *ethclient.Client

	token *Contract.BasToken
	miner *Contract.BasMiner
	asset *Contract.BasAsset
	oann  *Contract.BasOANN
	free  *Contract.SendFreeBas
)



func RestoreConnectionManually(entrance string) {
	c, err := ethclient.Dial(entrance)
	if err!=nil{
		logger.Error("can't connect to ethereum")
	}else{
		conn = c
		logger.Info("conn is ready")
	}
}

func GetConn() *ethclient.Client {
	if conn!=nil{
		return conn
	}else{
		for _,s := range AccessPoints{
			c, err := ethclient.Dial(s)
			if err!=nil{
				logger.Error("can't get access through : " ,s)
				continue
			}else{
				conn = c
				logger.Info("conn is ready")
				return conn
			}
		}
	}
	logger.Fatal("can't get access to ethereum through any points")
	return nil
}

func ResetConnAndContracts(){
	conn.Close()
	conn = nil
	token = nil
	miner = nil
	asset = nil
	oann = nil
	GetConn()
}

func GetLastBlockNumber() uint64{
	b,e:= GetConn().BlockByNumber(context.Background(),nil)
	if e==nil{
		return  b.NumberU64()
	}else {
		logger.Error("can't get last block",e)
		return 0
	}
}

func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(BASTokenAddress), GetConn());err==nil{
			token = t
		}else{
			logger.Error("can't recover BasToken",err)
		}
	}
	return token
}


func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(BASMinerAddress), GetConn());err==nil{
			miner = m
		}else{
			logger.Error("can't recover Miner",err)
		}
	}
	return miner
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(BASAssetAddress), GetConn());err==nil{
			asset = a
		}else{
			logger.Error("can't recover Asset",err)
		}
	}
	return asset
}

func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(BASOANNAddress), GetConn());err==nil{
			oann = o
		}else{
			logger.Error("can't recover OANN",err)
		}
	}
	return oann
}

func FreeBas() *Contract.SendFreeBas{
	if free==nil{
		if o,err:=Contract.NewSendFreeBas(common.HexToAddress(SendFreeBASAddress),GetConn());err==nil{
			free = o
		}else{
			logger.Error("can't recover Send Free Bas", err)
		}
	}
	return free
}
















