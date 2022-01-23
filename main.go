package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	io.WriteString(w, "'/post/{id}' - Find a post by its ID\n")
	io.WriteString(w, "'/post/{id}' Can also be used to delete a post!\n")

	fmt.Println("hit: '/'")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Posts)
	fmt.Println("hit: /posts")
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, post := range Posts {
		if post.Id == id {
			json.NewEncoder(w).Encode(post)
		}
	}
	fmt.Println("hit: '/posts/{id}' - by ID")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Blog
	json.Unmarshal(reqBody, &post)

	Posts = append(Posts, post)

	json.NewEncoder(w).Encode(Posts)

	fmt.Println("hit: '/posts/create'")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for i, post := range Posts {
		if post.Id == id {
			Posts = append(Posts[:i], Posts[i+1:]...)

		}
	}
	fmt.Println("hit: '/posts/{id}/delete'")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Welcome)
	myRouter.HandleFunc("/posts", GetAllPosts).Methods("GET")
	myRouter.HandleFunc("/post", CreatePost).Methods("POST")
	myRouter.HandleFunc("/post/{id}", GetPostByID).Methods("GET")
	myRouter.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")

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
		{
			Id:      "3",
			Title:   "PostThree",
			Author:  "Briana",
			Content: "Somethin about a drumkit, idk",
		},
		{
			Id:      "4",
			Title:   "PostFour",
			Author:  "Briana",
			Content: "Nutty Chocolate Cherry Trail Mix",
		},
		{
			Id:      "5",
			Title:   "PostFive",
			Author:  "Briana",
			Content: "Succotash does not belong in gumbo!",
		},
	}
	handleRequests()
}
