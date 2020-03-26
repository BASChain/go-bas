package main

import (
	"fmt"
	"github.com/BASChain/go-bas/Account"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/DataSync"
	"github.com/BASChain/go-bas/Market"
	"github.com/BASChain/go-bas/Transactions"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func main()  {
	testSync()
	//testMarket()
	//testCreatAccount()
	//Account.GetAuth("./key/UTC--2020-03-11T06-56-52.423772000Z--33324a5ee0b35f17536ceda27274e88e76640f24","secret")
	//testSendFreeEth()
	//testBalance()
	//fmt.Println(testCheckIfApplied())
	//testSendFreeBas()
	//fmt.Println(testCheckIfApplied())
}

func testSync(){
	DataSync.Sync();
	Bas_Ethereum.CloseConn()
}

func testMarket(){
	Market.Sync()
}

func testCreatAccount(){
	Account.CreateKs("secret");
}

func testResetConn(){
	Bas_Ethereum.BasAsset()
	Bas_Ethereum.ResetConnAndContracts()
	Bas_Ethereum.BasAsset()
}

func testSendFreeEth()  {
	keys:=Account.PrivateKeyRecover("./key/UTC--2020-03-11T06-56-52.423772000Z--33324a5ee0b35f17536ceda27274e88e76640f24","secret")
	key := keys[0]
	toAddress := common.HexToAddress("0x664ceB23bd9E35a246Fd20C31488604c8E470529")
	Transactions.SendFreeEth(key,toAddress,big.NewInt(1000000000000000))
}

func testBalance(){
	keys:=Account.PrivateKeyRecover("./key/UTC--2020-03-11T06-56-52.423772000Z--33324a5ee0b35f17536ceda27274e88e76640f24","secret")
	key := keys[0]
	fmt.Println(Transactions.CheckBalance(key).String())

}

func testCheckIfApplied()  (bool,error){
	return Transactions.CheckIfApplied(common.HexToAddress("0x5D34272bC26E5d6e475fbF25f63832b02f25Fb2a"))
}
func testSendFreeBas()  {
	keys:=Account.PrivateKeyRecover("./key/UTC--2020-03-11T06-56-52.423772000Z--33324a5ee0b35f17536ceda27274e88e76640f24","secret")
	key := keys[0]
	toAddress := common.HexToAddress("0x5D34272bC26E5d6e475fbF25f63832b02f25Fb2a")
	Transactions.SendFreeBasByContract(key,toAddress,big.NewInt(1000000000000000))
}