package Transactions

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/op/go-logging"
	"math/big"
)


var logger, _ = logging.GetLogger("Transactions")

func SendFreeEth(key *keystore.Key,toAddress common.Address,amount int64) {

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := Bas_Ethereum.GetConn().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Error("send free eth, create none error : ",err)
		return
	}

	value := big.NewInt(amount) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := Bas_Ethereum.GetConn().SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("get gasPrice error : ", err)
		return
	}

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := Bas_Ethereum.GetConn().NetworkID(context.Background())
	if err != nil {
		logger.Error("get chainId error : ", err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		logger.Error("sign transaction error : ",err)
		return
	}

	err = Bas_Ethereum.GetConn().SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Error("failed to send transaction : ", err)
	}

	receipt, err := bind.WaitMined(context.Background(),Bas_Ethereum.GetConn(),signedTx)

	if err != nil {
		logger.Error("send free bas error :", err)
		return
	}else{
		fmt.Println("send free bas to : ",toAddress.String()," on block : ",receipt.BlockNumber.String())
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}