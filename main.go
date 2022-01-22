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

	fmt.Println("hit: '/'")
}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, post := range Posts {
		if post.Id == key {
			json.NewEncoder(w).Encode(post)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Welcome)
	myRouter.HandleFunc("/Posts", GetAllBlogs)
	myRouter.HandleFunc("/Post/{id}", GetBlogByID)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	Posts = []Blog{
		{
			Title:   "PostOne",
			Author:  "Spencer",
			Id:      "1",
			Content: "Once upon a time...",
		},
		{
			Title:   "PostTwo",
			Author:  "Spencer",
			Id:      "2",
			Content: "This is the content of the second post",
		},
	}

	handleRequests()
}
