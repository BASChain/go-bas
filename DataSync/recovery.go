package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("DataSync")

var (

	TokenAddr     = "0x9d0314f9Bacd569DCB22276867AAEeE1C8A87614"
	OwnershipAddr = "0x4b91b82bed39B1d946C9E3BC12ba09C2F22fd3ee"
	AssetAddr     = "0x2B1110a13183A7045C7BCE3ba0092Ff0de4FD241"
	DNSAddr       = "0x8951f6B80b880E8A47d0d18000A4c90F288F61a3"
	OANNAddr      = "0x540d08bB1B16e09E1AefD56147c9435f83BBd437"

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
		if t,err:=Contract.NewBasToken(common.HexToAddress(TokenAddr), conn.GetClient());err==nil{
			token = t
		}else{
			logger.Fatal("can't recover BasToken",err)
		}
	}
	return token
}

func BasOwnership() *Contract.BasOwnership  {
	if ownership==nil{
		if o,err:=Contract.NewBasOwnership(common.HexToAddress(OwnershipAddr), conn.GetClient());err==nil{
			ownership = o
		}else{
			logger.Fatal("can't recover ownership",err)
		}
	}
	return ownership
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(AssetAddr), conn.GetClient());err==nil{
			asset = a
		}else{
			logger.Fatal("can't recover Asset",err)
		}
	}
	return asset
}

func BasDNS() *Contract.BasDNS{
	if dns==nil{
		if d,err:=Contract.NewBasDNS(common.HexToAddress(DNSAddr), conn.GetClient());err==nil{
			dns = d
		}else{
			logger.Fatal("can't recover dns",err)
		}
	}
	return dns
}


func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(OANNAddr), conn.GetClient());err==nil{
			oann = o
		}else{
			logger.Fatal("can't recover OANN",err)
		}
	}
	return oann
}
















