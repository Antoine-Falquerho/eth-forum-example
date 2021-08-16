package main

import (
	"text/template"
	"fmt"
	"github.com/Antoine-Falquerho/eth-forum/helpers"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/schema"
	"errors"
)

// https://towardsdev.com/creating-a-simple-ethereum-smart-contract-in-golang-138b9439f64e
func main(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/users/update", updateAccount)
	http.HandleFunc("/posts/new", addPostRequest)
	//http.HandleFunc("/view/", makeHandler(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//https://freshman.tech/web-development-with-go/
func homePage(w http.ResponseWriter, r *http.Request){
	posts := helpers.GetPosts()
	fmt.Println("Endpoint Hit: homePage")
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

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
// To use cookies
func login(w http.ResponseWriter, r *http.Request){
	var p = make(map[string]string)
	json.NewDecoder(r.Body).Decode(&p)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "address", Value: p["address"], Expires: expiration}
	http.SetCookie(w, &cookie)
}


type postForm struct {
	Title string
	Content string
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

type UserForm struct {
	Username string
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

func getSessionAddress(r *http.Request) (string, error){
	userAddress, err := r.Cookie("address")
	if err != nil {
		return "", errors.New("not found")
	}
	return userAddress.Value, nil
}

