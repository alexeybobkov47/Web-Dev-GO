package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (database *Server) showBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := getBlogs(database.db)
	if err != nil {
		log.Println(err)
		return
	}

	if err := tmplBlog.ExecuteTemplate(w, "blog", blogs); err != nil {
		log.Println(err)

	}

}

func (database *Server) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	p, err := strconv.Atoi(path[2])
	if err != nil {
		log.Println(err)
		return
	}
	posts, err := getPosts(database.db, (p - 1))
	if err := tmplPost.ExecuteTemplate(w, "post", posts); err != nil {
		log.Println(err)

	}

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
