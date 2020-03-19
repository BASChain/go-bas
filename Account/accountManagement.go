package Account

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/op/go-logging"
	"os"
	"strings"
)

var logger, _ = logging.GetLogger("Account")

func CreateKs(password string) {
	ks := keystore.NewKeyStore("./key", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		logger.Error("create account err : ",err)
	}

	fmt.Println(account.Address.Hex())
}

func GetAuth(keyFile string, pass string) []*bind.TransactOpts {
	keys:= ReadTextByline(keyFile)
	var auth  []*bind.TransactOpts
	for i:=0;i<len(keys);i++{
		a, _ := bind.NewTransactor(strings.NewReader(keys[i]), pass)
		auth = append(auth, a)
	}
	return auth
}

func PrivateKeyRecover(keyFile string, pass string)[]*keystore.Key{
	keys:= ReadTextByline(keyFile)
	var keyList  []*keystore.Key
	for i:=0;i<len(keys);i++{
		key, err := keystore.DecryptKey([]byte(keys[i]), pass)
		if err == nil{
			keyList = append(keyList, key)
		}
	}
	return keyList
}

func PrivateKeyRecoverByBytes(privData []byte,pass string) []*keystore.Key {
	keys := strings.Split(string(privData),"\n")

	var keyList  []*keystore.Key
	for i:=0;i<len(keys);i++{
		key, err := keystore.DecryptKey([]byte(keys[i]), pass)
		if err == nil{
			keyList = append(keyList, key)
		}
	}
	return keyList

}


func ReadTextByline(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("read file error : ", err)
	}
	defer file.Close()
	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys,scanner.Text() )
	}

	if err := scanner.Err(); err != nil {
		logger.Error("read line error : ",err)
	}
	return keys
}


