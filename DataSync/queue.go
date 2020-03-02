package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
)

var queueAsset []Bas_Ethereum.Hash
var queueDns []Bas_Ethereum.Hash


func insertQueueAsset(key Bas_Ethereum.Hash){
	if !exists(queueAsset,key) {
		queueAsset = append(queueAsset, key)
	}
}


func clearQueryAsset(){
	queueAsset = []Bas_Ethereum.Hash{}
}

func insertQueueDns(key Bas_Ethereum.Hash){
	if !exists(queueDns,key){
		queueDns = append(queueDns, key)
	}

}

func clearQueueDns(){
	queueDns = []Bas_Ethereum.Hash{}
}