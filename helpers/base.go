package helpers

import (
	"fmt"
	"github.com/Antoine-Falquerho/eth-forum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	CONTRACT_ADDRESS = "0xC858EC06fEa37A1A6255F6E299e0F3165D19828d"
	WALLET_ADDRESS = "0x19b2b877119EeB019C0E03F24655c12C3f1b4288"
	PRIVATE_KEY = "55b9a9e94f0a01f9e6cead2a42887f68bd674e8f28a5cef6295d6631fb8537d0"
)

func getConn() *contracts.Forum{
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	conn, err := contracts.NewForum(contractAddress, client)
	if err != nil {
		panic(err)
	}
	return conn
}

func getTransactOpts() (*bind.TransactOpts ){
	fmt.Println("Hello World")

	//walletAddress := common.HexToAddress(WALLET_ADDRESS)
	privateKey, _ := crypto.HexToECDSA(PRIVATE_KEY)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatal("error getTransactOpts")
	}
	return transactOpts
}