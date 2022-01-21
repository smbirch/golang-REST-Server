package main

import (
	"log"
	"net/http"
	"time"
)

//storage for blogs

//The following methods will work from some type of local, non DB datastore
//
//CreateBlog() creates a new blog post - POST /blogs
//
//GetBlog() will retrieve the blog post - GET /blogs/{Id}
//
//DeleteBlog() will delete the blog post - DELETE /blogs/{Id}
//
//GetAllBlogs() will return all blog posts - GET /blogs
//
//GetBlogByDate() will return all blog posts from a given Date

//

//blog type
type Blog struct {
	Title   string    `json:"title"`   //blog title
	Author  string    `json:"author"`  //author name
	Id      int64     `json:"id"`      //id number of blog post
	Content string    `json:"content"` //content of blog post ie text
	Date    time.Time `json:"date"`    //the date this post was created
}

type blogHandlers struct {
	store map[int]Blog
}

func (b *blogHandlers) getBlog(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/blogs", blogsHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//The word blog has lost all meaning
