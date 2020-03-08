package main

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/DataSync"
)

func main()  {

	//testResetConn()
	DataSync.Sync();
}

func testResetConn(){
	Bas_Ethereum.BasAsset()
	Bas_Ethereum.ResetConnAndContracts()
	Bas_Ethereum.BasAsset()
}
