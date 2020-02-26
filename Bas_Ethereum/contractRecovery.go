package Bas_Ethereum

import (
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"sync"
)
var logger, _ = logging.GetLogger("Bas_Ethereum")

var (
	AccessPoint = "https://ropsten.infura.io/v3/8b8db3cca50a4fcf97173b7619b1c4c3"

	BASTokenAddress = "0xb9291744e91fAd68060D7Cbfc702f13F64E6E7C5"
	BASMinerAddress = "0x3C81935818a7C8e08372C3aE457AD5F8719B9901"
	BASAssetAddress = "0xFD2d0B61f1d956CFC69b73c60F7647f3a1b9500D"
	BASOANNAddress = "0xCf1FFcFB1A6e1ADfde8bBd7B59636D5cB4080355"
)

var recoverLock sync.Mutex
var conn *ethclient.Client

var Token *Contract.BasToken
var Miner *Contract.BasMiner
var Asset *Contract.BasAsset
var OANN  *Contract.BasOANN

func resetAccessPoint(entrance string) {
	AccessPoint = entrance
}

func getConn() *ethclient.Client {
	if conn!=nil{
		return conn
	}else{
		c, err := ethclient.Dial(AccessPoint)
		if err!=nil{
			logger.Error("can't get access to ethereum",err)
			return nil;
		}else{
			conn = c
			return conn
		}
	}
}


func RecoverBasToken() *Contract.BasToken{
	if token,err:=Contract.NewBasToken(common.HexToAddress(BASTokenAddress),getConn());err==nil{
		Token = token
	}else{
		logger.Error("can't recover BasToken",err)
	}
	return Token
}


func RecoverBasMiner() *Contract.BasMiner{
	if miner,err:=Contract.NewBasMiner(common.HexToAddress(BASMinerAddress),getConn());err==nil{
		Miner = miner
	}else{
		logger.Error("can't recover Miner",err)
	}
	return Miner
}

func RecoverBasAsset() *Contract.BasAsset{
	if asset,err:=Contract.NewBasAsset(common.HexToAddress(BASAssetAddress),getConn());err==nil{
		Asset = asset
	}else{
		logger.Error("can't recover Asset",err)
	}
	return Asset
}

func RecoverBasOANN() *Contract.BasOANN{
	if oann,err:=Contract.NewBasOANN(common.HexToAddress(BASOANNAddress),getConn());err==nil{
		OANN = oann
	}else{
		logger.Error("can't recover OANN",err)
	}
	return OANN
}


















