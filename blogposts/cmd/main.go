package main

import (
	"log"
	"os"

	blogpost "github.com/tumypmyp/learn-go-with-tests/blogposts"
)

func main() {
	posts, err := blogpost.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("posts: %+v", posts)
}
