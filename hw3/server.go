package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var blog1 = Blog{
	Name:        "Личный блог",
	Description: "Посты на разные темы",
	Posts: []Post{
		{ID: "1", Header: "Первый пост", Text: "13241 3g r1344 23 3fdgdg "},
		{ID: "2", Header: "Второй пост", Text: "hjghj 45456 435th hhth4 hth"},
		{ID: "3", Header: "Третий Пост", Text: "th4th4 t4ht 4th th 4tht ht"},
	},
}

var tmplBlog = template.Must(template.New("BlogTemplate").ParseFiles("index.html"))
var tmplPost = template.Must(template.New("PostTemplate").ParseFiles("post.html"))
var tmplNewPost = template.Must(template.New("PostTemplate").ParseFiles("newpost.html"))

func main() {
	router := http.NewServeMux()
	port := "8080"

	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("src"))))
	router.HandleFunc("/blog", showBlog)
	for _, value := range blog1.Posts {
		log.Println(value.ID)
		router.HandleFunc("/blog/"+value.ID, showPost)
	}
	router.HandleFunc("/blog/newpost", newPost)

	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func showBlog(w http.ResponseWriter, r *http.Request) {
	if err := tmplBlog.ExecuteTemplate(w, "blog", blog1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showPost(w http.ResponseWriter, r *http.Request) {
	path, err := strconv.Atoi(r.URL.Path[6:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmplPost.ExecuteTemplate(w, "post", blog1.Posts[path-1]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func newPost(w http.ResponseWriter, r *http.Request) {

	newpost := Post{
		ID:     r.FormValue("ID"),
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}

	if newpost.ID != "" {
		blog1.Posts = append(blog1.Posts, newpost)
	}

	if err := tmplNewPost.ExecuteTemplate(w, "newpost", blog1.Posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}