package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Blog struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var Posts []Blog

func Welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the blog!\n\n")
	io.WriteString(w, "'/posts' - See all posts\n")
	io.WriteString(w, "'/posts/{id}' - Find a post by its ID\n")
	io.WriteString(w, "'/posts/{id}' Can also be used to delete a post!")

	fmt.Println("hit: '/'")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Posts)
	fmt.Println("hit: /posts")
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, post := range Posts {
		if post.Id == key {
			json.NewEncoder(w).Encode(post)
		}
	}
	fmt.Println("hit: '/posts/{id}'")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Welcome)
	myRouter.HandleFunc("/posts", GetAllPosts)
	myRouter.HandleFunc("/posts/{id}", GetPostByID)
	myRouter.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	Posts = []Blog{
		{
			Id:      "1",
			Title:   "PostOne",
			Author:  "Spencer",
			Content: "Once upon a time...",
		},
		{
			Id:      "2",
			Title:   "PostTwo",
			Author:  "Spencer",
			Content: "This is the content of the second post",
		},
	}
	handleRequests()
}
