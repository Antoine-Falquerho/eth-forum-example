package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type User struct {
	Username string
	Karma big.Int
	PostCount big.Int
}

func GetUser(walletAddress string) User{
	conWalletAddress := common.HexToAddress(walletAddress)
	conn := getConn()

	user, err := conn.GetUser(&bind.CallOpts{Pending: true}, conWalletAddress)
	if err != nil {
		panic(err)
	}

	return User{user.Name, *user.Karma, *user.PostCount}
}


func UpdateName(walletAddress string, newName string){
	conWalletAddress := common.HexToAddress(walletAddress)
	conn := getConn()
	transactOpts := getTransactOpts()
	_, err := conn.SetName(transactOpts, conWalletAddress, newName)
	if err != nil {
		panic(err)
	}
}
