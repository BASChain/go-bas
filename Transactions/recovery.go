package Transactions

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	_f = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"
)

func FreeBas() (*Contract.SendFreeBas,*ethclient.Client){

	conn:=Bas_Ethereum.NewConn()
	//no watch, always get a new conn
	if o,err:=Contract.NewSendFreeBas(common.HexToAddress(_f),conn.GetClient());err==nil{
		return o,conn.Client
	}else{
		logger.Error("can't recover Send Free Bas", err)
		return nil,nil
	}
}