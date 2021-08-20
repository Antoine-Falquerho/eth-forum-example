package helpers

import (
	"fmt"
	"log"
	"time"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Post struct {
	ID int64
	Owner common.Address
	Title string
	Content string
	Karma int64
	PostedAt string
	User User
}

func GetPosts() []Post{
	conn := getConn()
	posts := []Post{}
	lastPostId, _ := conn.LastPostId(&bind.CallOpts{Pending: true})
	for i := lastPostId.Int64(); i > 0; i = i -1{
		post_id := i -1
		post, _ := conn.Posts(&bind.CallOpts{Pending: true}, big.NewInt(post_id))
		user := GetUser(post.Owner.String())
		postedAt := time.Unix(post.PostedAt.Int64(), 0)
		posts = append(
			posts,
			Post{
				post_id,
				post.Owner,
				post.Title,
				post.Content,
				post.Karma.Int64(),
				postedAt.Format("02/01/2006, 15:04:05"),
				user,
			},
		)
	}
	return posts
}

func AddPost(walletAddress string, title string, content string){
	conWalletAddress := common.HexToAddress(walletAddress)
	conn := getConn()
	transactOpts := getTransactOpts()

	_, err := conn.AddPost(transactOpts, conWalletAddress, title, content)
	if err != nil {
		log.Fatal(err)
	}
}

func AddVote(address string, post_id *big.Int, vote *big.Int){
	conWalletAddress := common.HexToAddress(address)
	conn := getConn()
	transactOpts := getTransactOpts()
	_, err := conn.AddVote(transactOpts, conWalletAddress, post_id, vote)
	if err != nil {
		log.Fatal(err)
	}
}
