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
	AccessPoint = "wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3"

	BASTokenAddress = "0xb9291744e91fAd68060D7Cbfc702f13F64E6E7C5"
	BASMinerAddress = "0x3C81935818a7C8e08372C3aE457AD5F8719B9901"
	BASAssetAddress = "0xCDF7A205DE6f35b60a540dE5B9ebF419c2303e61"
	BASOANNAddress = "0x53CAb752Fd2c5570CBb3F013Ea034d2AC45Dc6C7"

	conn *ethclient.Client

	token *Contract.BasToken
	miner *Contract.BasMiner
	asset *Contract.BasAsset
	oann  *Contract.BasOANN
)



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

func GetLastBlockNumber() uint64{
	b,e:=getConn().BlockByNumber(context.Background(),nil)
	if e==nil{
		return  b.NumberU64()
	}else {
		logger.Error("can't get last block",e)
		return 0
	}
}

func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(BASTokenAddress),getConn());err==nil{
			token = t
		}else{
			logger.Error("can't recover BasToken",err)
		}
	}
	return token
}


func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(BASMinerAddress),getConn());err==nil{
			miner = m
		}else{
			logger.Error("can't recover Miner",err)
		}
	}
	return miner
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(BASAssetAddress),getConn());err==nil{
			asset = a
		}else{
			logger.Error("can't recover Asset",err)
		}
	}
	return asset
}

func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(BASOANNAddress),getConn());err==nil{
			oann = o
		}else{
			logger.Error("can't recover OANN",err)
		}
	}
	return oann
}


















