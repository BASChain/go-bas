package Transactions

// this file should not be included in product environment!

import (
	"context"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var (
	FreeBasAddr = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"
)

func FreeBas() (*Contract.SendFreeBas,*ethclient.Client){
	conn:=Bas_Ethereum.NewConn()
	//no watch, always get a new conn
	if o,err:=Contract.NewSendFreeBas(common.HexToAddress(FreeBasAddr),conn.GetClient());err==nil{
		return o,conn.Client
	}else{
		logger.Error("can't recover Send Free Bas", err)
		return nil,nil
	}
}

func CheckIfApplied(addr common.Address) (bool,error){
	free,client:=FreeBas()
	defer client.Close()
	return free.ApplyRecord(nil,addr)
}


func SendFreeBasByContract(key *keystore.Key,addr common.Address,amount *big.Int) error {
	opts:=bind.NewKeyedTransactor(key.PrivateKey)
	free,client:=FreeBas()
	defer client.Close()
	tx,err:=free.SendTokenByContract(opts,addr,amount)
	if err!=nil{
		logger.Error("send BAS error : ",err)
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err!=nil{
		logger.Error("mine send BAS error : ", err)
		return err
	}else{
		logger.Info("send free bas to : ",addr.String()," on block : ",receipt.BlockNumber.String())
	}
	return nil
}