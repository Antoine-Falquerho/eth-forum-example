package main

import (
	"text/template"
	"fmt"
	"github.com/Antoine-Falquerho/eth-forum/helpers"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/Antoine-Falquerho/eth-forum/contracts"
	"log"
	"net/http"
)





// https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e
func main(){
	//client, err := ethclient.Dial("http://127.0.0.1:7545")
	//if err != nil {
	//	panic(err)
	//}
	//contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	//conn, err := contracts.NewForum(contractAddress, client)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Hello World")
	//
	//walletAddress := common.HexToAddress(WALLET_ADDRESS)
	//privateKey, _ := crypto.HexToECDSA(PRIVATE_KEY)
	//transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	//u, err := conn.SetName(transactOpts, walletAddress, "Proutt")
	//fmt.Println(u)
	//fmt.Println(err)
	//fmt.Println("--- -_____---- ----")
	//fmt.Println(conn.LastPostId(&bind.CallOpts{Pending: true}))

	//fmt.Println(err)
	//if u != nil {
	//	fmt.Println("good")
	//}

	//addPost(WALLET_ADDRESS, "Antoine", "Flqh")

	//user, err := conn.GetUser(&bind.CallOpts{Pending: true}, walletAddress)
	//fmt.Println(user.Name)
	//fmt.Println(user.Karma)
	//fmt.Println(user.PostCount)

	//handleRequests()
	// creates a new instance of a mux router
	// creates a new instance of a mux router

	http.HandleFunc("/", homePage)
	//http.HandleFunc("/view/", makeHandler(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//https://freshman.tech/web-development-with-go/
func homePage(w http.ResponseWriter, r *http.Request){
	//updateName(WALLET_ADDRESS, "Bob")
	//addPost(WALLET_ADDRESS, "Kim", "Steele")
	posts := helpers.GetPosts()
	fmt.Println("Endpoint Hit: homePage")
	user := helpers.GetUser(helpers.WALLET_ADDRESS)
	type response struct {
		User helpers.User
		Posts []helpers.Post
	}
	p := response{
		User: user,
		Posts: posts,
	}

	//t, _ := template.ParseFiles("template/index.gohtml")
	t := template.Must(template.ParseFiles("template/index.gohtml"))
	t.Execute(w, p)
}









