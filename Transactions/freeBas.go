package Transactions

import (
	"context"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func CheckIfApplied(addr common.Address) (bool,error){
	return Bas_Ethereum.FreeBas().ApplyRecord(nil,addr)
}


func SendFreeBasByContract(key *keystore.Key,addr common.Address,amount *big.Int) error {
	opts:=bind.NewKeyedTransactor(key.PrivateKey)
	tx,err:=Bas_Ethereum.FreeBas().SendTokenByContract(opts,addr,amount)
	if err!=nil{
		logger.Error("send BAS error : ",err)
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), Bas_Ethereum.GetConn(), tx)
	if err!=nil{
		logger.Error("mine send BAS error : ", err)
		return err
	}else{
		logger.Info("send free bas to : ",addr.String()," on block : ",receipt.BlockNumber.String())
	}
	return nil
}

