package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Blog struct {
	Title   string    `json:"title"`   //blog title
	Author  string    `json:"author"`  //author name
	Id      string    `json:"id"`      //id number of blog post
	Content string    `json:"content"` //content of blog post ie text
	Date    time.Time `json:"date"`    //the date this post was created
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the blog!\n")
}

func BlogsHandler(w http.ResponseWriter, r *http.Request) {

	blogDB := []Blog{
		{
			Title:   "PostOne",
			Author:  "Spencer",
			Id:      "1",
			Content: "Once upon a time...",
			Date:    time.Now(),
		},
		{
			Title:   "PostTwo",
			Author:  "Spencer",
			Id:      "2",
			Content: "This is the content of the second post",
			Date:    time.Now(),
		},
	}
	json.NewEncoder(w).Encode(blogDB)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)

	mux.HandleFunc("/blogs", BlogsHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
