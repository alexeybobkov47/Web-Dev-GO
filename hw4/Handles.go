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
	}

	if err := tmplBlog.ExecuteTemplate(w, "blog", blogs); err != nil {
		log.Println(err)
	}
	return
}

func (database *Server) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	p := strconv.Itoa(len(path) - 1)
	posts, err := getPosts(database.db, p)
	if err != nil {
		log.Println(err)
	}
	if err := tmplPost.ExecuteTemplate(w, "post", posts); err != nil {
		log.Println(err)

	}
	return
}

func (database *Server) newPost(w http.ResponseWriter, r *http.Request) {
	newpost := Post{
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}

	if len(newpost.Header) != 0 && len(newpost.Text) != 0 {
		newPost(database.db, newpost)
	}

	if err := tmplNewPost.ExecuteTemplate(w, "newpost", newpost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
