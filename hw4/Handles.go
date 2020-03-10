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
		return
	}

	if err := tmplBlog.ExecuteTemplate(w, "blog", blogs); err != nil {
		log.Println(err)
		return
	}
}

func (database *Server) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	posts, err := getPosts(database.db, (path[len(path)-1]))
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmplPost.ExecuteTemplate(w, "post", posts); err != nil {
		log.Println(err)
		return
	}
}

func (database *Server) newPost(w http.ResponseWriter, r *http.Request) {
	newpost := Post{
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}

	if len(newpost.Header) != 0 && len(newpost.Text) != 0 {
		err := newPost(database.db, newpost)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err := tmplNewPost.ExecuteTemplate(w, "newpost", newpost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (database *Server) editPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	p := (path[len(path)-1])
	posts, err := getPosts(database.db, p)
	if err != nil {
		log.Println(err)
		return
	}

	editpost := Post{
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}
	if len(editpost.Header) != 0 && len(editpost.Text) != 0 {
		err := editPost(database.db, editpost, p)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err := tmplEditPost.ExecuteTemplate(w, "editpost", posts); err != nil {
		log.Println(err)
		return
	}
}
