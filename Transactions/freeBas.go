package Transactions

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)



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

