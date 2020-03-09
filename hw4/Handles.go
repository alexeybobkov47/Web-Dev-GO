package main

import (
	"log"
	"net/http"
	"strings"
)

func (database *Server) showBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := getBlogs(database.db)
	if err != nil {
		log.Println(err)
	}

	if err := tmplBlog.ExecuteTemplate(w, "blog", blogs); err != nil {
		log.Println(err)
	}
	return
}

func (database *Server) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	posts, err := getPosts(database.db, path[2])
	if err != nil {
		log.Println(err)
	}
	if err := tmplPost.ExecuteTemplate(w, "post", posts); err != nil {
		log.Println(err)

	}
	return
}

// func (database *sql.DB) newPost(w http.ResponseWriter, r *http.Request) {
// 	lenBlog := len(blog1.Posts)
// 	newpost := Post{
// 		ID:     lenBlog + 1,
// 		Header: r.FormValue("header"),
// 		Text:   r.FormValue("text"),
// 	}

// 	if len(newpost.Header) != 0 {
// 		blog1.Posts = append(blog1.Posts, newpost)
// 	}

// 	if err := tmplNewPost.ExecuteTemplate(w, "newpost", blog1.Posts); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }
