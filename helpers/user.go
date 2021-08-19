package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	Username string
	Karma string
	PostCount string
	Address string
}

func GetUser(walletAddress string) User{
	conWalletAddress := common.HexToAddress(walletAddress)
	conn := getConn()

	user, err := conn.GetUser(&bind.CallOpts{Pending: true}, conWalletAddress)
	if err != nil {
		panic(err)
	}
	return User{user.Name, (*user.Karma).String(), (*user.PostCount).String(), walletAddress}
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
