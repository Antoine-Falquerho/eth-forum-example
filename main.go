package main

import (
	"text/template"
	"github.com/Antoine-Falquerho/eth-forum/helpers"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/schema"
	"math/big"
	"errors"
)

type postForm struct {
	Title string
	Content string
}

type UserForm struct {
	Username string
}

// https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e
func main(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/users/update", updateAccount)
	http.HandleFunc("/posts/new", addPostRequest)
	http.HandleFunc("/add_user_vote", addUserVote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//https://freshman.tech/web-development-with-go/
func homePage(w http.ResponseWriter, r *http.Request){
	posts := helpers.GetPosts()
	address, err := getSessionAddress(r)
	if err != nil{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	user := helpers.GetUser(address)
	type response struct {
		User helpers.User
		Posts []helpers.Post
	}
	p := response{
		User: user,
		Posts: posts,
	}

	t := template.Must(template.ParseFiles("template/index.gohtml"))
	t.Execute(w, p)
}

func login(w http.ResponseWriter, r *http.Request){
	var p = make(map[string]string)
	json.NewDecoder(r.Body).Decode(&p)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "address", Value: p["address"], Expires: expiration}
	http.SetCookie(w, &cookie)
	w.Write([]byte("OK"))
}

func addPostRequest(w http.ResponseWriter, r *http.Request){
	var newPost = postForm{}
	if err := r.ParseForm(); err != nil {
		// handle error
	}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&newPost, r.PostForm)
	if err != nil {
		// Handle error
	}

	address, err := getSessionAddress(r)
	if err != nil{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	helpers.AddPost(address, newPost.Title, newPost.Content)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateAccount(w http.ResponseWriter, r *http.Request){
	var updateUser = UserForm{}
	if err := r.ParseForm(); err != nil {
		// handle error
	}
	decoder := schema.NewDecoder()
	err := decoder.Decode(&updateUser, r.PostForm)
	if err != nil {
		// Handle error
	}
	address, err := getSessionAddress(r)
	if err != nil{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	helpers.UpdateName(address, updateUser.Username)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addUserVote(w http.ResponseWriter, r *http.Request){
	vote, ok := r.URL.Query()["vote"]
	if !ok || len(vote[0]) < 1 {
		log.Println("Url Param 'vote' is missing")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	post_id := new(big.Int)
	post_id, ok = post_id.SetString(r.URL.Query()["post_id"][0], 10)

	address, err := getSessionAddress(r)
	if err != nil{
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	vote_int := 1
	if vote[0] == "down" {
		vote_int = -1
	}
	helpers.AddVote(address, post_id, big.NewInt(int64(vote_int)))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getSessionAddress(r *http.Request) (string, error){
	userAddress, err := r.Cookie("address")
	if err != nil {
		return "", errors.New("not found")
	}
	return userAddress.Value, nil
}

