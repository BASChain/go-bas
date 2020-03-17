package Transactions

import (
	"context"
	"crypto/ecdsa"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/op/go-logging"
	"math/big"
	"errors"
	"github.com/BASChain/go-bas-dns-server/dns/mem"
)


var logger, _ = logging.GetLogger("Transactions")

func CheckBalance(key *keystore.Key) *big.Int{
	account := GetPublicAddressFromKeyStore(key)
	balance, err := Bas_Ethereum.GetConn().BalanceAt(context.Background(), account, nil)
	if err != nil {
		logger.Error("check balance failed",err)
	}
	return  balance
}

func GetPublicAddressFromKeyStore(key *keystore.Key) common.Address {
	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return  crypto.PubkeyToAddress(*publicKeyECDSA)
}

func SendFreeEth(key *keystore.Key,toAddress common.Address,amount *big.Int) error {
	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("error casting public key to ECDSA")
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := Bas_Ethereum.GetConn().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Error("send free eth, create none error : ",err)
		return err
	}

	gasLimit := uint64(21000)                // in units
	gasPrice, err := Bas_Ethereum.GetConn().SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("get gasPrice error : ", err)
		return err
	}

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	chainID, err := Bas_Ethereum.GetConn().NetworkID(context.Background())
	if err != nil {
		logger.Error("get chainId error : ", err)
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		logger.Error("sign transaction error : ",err)
		return err
	}

	err = Bas_Ethereum.GetConn().SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Error("failed to send transaction : ", err)
	}

	receipt, err := bind.WaitMined(context.Background(),Bas_Ethereum.GetConn(),signedTx)

	if err != nil {
		logger.Error("send free eth error :", err)
		return err
	}else{
		logger.Info("send free eth to : ",toAddress.String()," on block : ",receipt.BlockNumber.String())
	}

	return nil
}

func SendFreeEthWrapper(key *keystore.Key,toAddress common.Address,amount *big.Int)  {
	mem.Update(toAddress,mem.ETH,mem.WAITING)
	go func() {
		if err:=SendFreeEth(key,toAddress,amount);err!=nil{
			mem.Update(toAddress,mem.ETH,mem.FAILURE)
		}else{
			mem.Update(toAddress,mem.ETH,mem.SUCCESS)
		}
	}()
}