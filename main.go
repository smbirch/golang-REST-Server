package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//storage for blogs

//The following methods will work from some type of local, non DB datastore
//
//CreateBlog() creates a new blog post - POST /blogs
//
//GetBlog() will retrieve the blog post - GET /blogs/{Id} - WIP
//
//DeleteBlog() will delete the blog post - DELETE /blogs/{Id}
//
//GetAllBlogs() will return all blog posts - GET /blogs - Done?!
//
//GetBlogByDate() will return all blog posts from a given Date

//

//blog type
type Blog struct {
	Title   string    `json:"title"`   //blog title
	Author  string    `json:"author"`  //author name
	Id      int       `json:"id"`      //id number of blog post
	Content string    `json:"content"` //content of blog post ie text
	Date    time.Time `json:"date"`    //the date this post was created
}

type blogsHandler struct {
	store map[int]Blog //map of Id key
}

func (b *blogsHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := make([]Blog, len(b.store))

	i := 0
	for _, blog := range b.store {
		blogs[i] = blog
		i++
	}
	jsonBytes, err := json.Marshal(blogs)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonBytes)
}

//constructor returning pointer to a new blog handler
func newBlogsHandler() *blogsHandler {
	return &blogsHandler{
		store: map[int]Blog{
			1: {
				Title:   "PostOne",
				Author:  "Spencer",
				Id:      1,
				Content: "Once upon a time...",
				Date:    time.Now(),
			},
			2: {
				Title:   "PostTwo",
				Author:  "Spencer",
				Id:      2,
				Content: "This is the content of the second post",
				Date:    time.Now(),
			},
		},
	}

}

func main() {

	blogsHandler := newBlogsHandler()

	http.HandleFunc("/blogs", blogsHandler.GetAllBlogs)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//The word blog has lost all meaning
