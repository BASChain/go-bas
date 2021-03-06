package DataSync

import "github.com/BASChain/go-bas/Bas_Ethereum"

func deleteFromStringList(list []Bas_Ethereum.Hash, key Bas_Ethereum.Hash) []Bas_Ethereum.Hash{
	if list == nil{
		return  nil
	}
	for i,k := range list {
		if k==key {
			return append(list[:i],list[i+1:]...)
		}
	}
	return list
}

func exists(list []Bas_Ethereum.Hash, key Bas_Ethereum.Hash) bool {
	if list == nil{
		return false
	}
	for _,k :=range list{
		if k==key {
			return true
		}
	}
	return false
}