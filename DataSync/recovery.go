package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
)

var (

	_t = "0x9d0314f9Bacd569DCB22276867AAEeE1C8A87614"
	_o = "0xd6bA5488a46028e3346ed792c4E74f67c0e11DD4"
	_a = "0x2B1110a13183A7045C7BCE3ba0092Ff0de4FD241"
	_d = "0x8951f6B80b880E8A47d0d18000A4c90F288F61a3"
	_oann = "0xeC2218E1e40CB4b21a4142Df6ee8a6eAF4Bb4434"

	conn = Bas_Ethereum.NewConn()

	token *Contract.BasToken
	ownership *Contract.BasOwnership
	asset *Contract.BasAsset
	dns *Contract.BasDNS
	oann  *Contract.BasOANN
)





func ResetConnAndService(){
	conn.Reset()
	token = nil
	ownership = nil
	asset = nil
	dns = nil
	oann = nil
}


func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(_t), conn.GetClient());err==nil{
			token = t
		}else{
			logger.Fatal("can't recover BasToken",err)
		}
	}
	return token
}

func BasOwnership() *Contract.BasOwnership  {
	if ownership==nil{
		if o,err:=Contract.NewBasOwnership(common.HexToAddress(_o), conn.GetClient());err==nil{
			ownership = o
		}else{
			logger.Fatal("can't recover ownership",err)
		}
	}
	return ownership
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(_a), conn.GetClient());err==nil{
			asset = a
		}else{
			logger.Fatal("can't recover Asset",err)
		}
	}
	return asset
}

func BasDNS() *Contract.BasDNS{
	if dns==nil{
		if d,err:=Contract.NewBasDNS(common.HexToAddress(_d), conn.GetClient());err==nil{
			dns = d
		}else{
			logger.Fatal("can't recover dns",err)
		}
	}
	return dns
}


func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(_oann), conn.GetClient());err==nil{
			oann = o
		}else{
			logger.Fatal("can't recover OANN",err)
		}
	}
	return oann
}
















