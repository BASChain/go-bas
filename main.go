package main

import "github.com/BASChain/go-bas/Bas_Ethereum"

func main()  {
	Bas_Ethereum.RecoverBasAsset();
	data,_:=Bas_Ethereum.QueryDomain("root1")
	print(data.AName)
}
