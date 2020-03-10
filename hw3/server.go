package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	tmplBlog    = template.Must(template.New("BlogTemplate").ParseFiles("index.html"))
	tmplPost    = template.Must(template.New("PostTemplate").ParseFiles("post.html"))
	tmplNewPost = template.Must(template.New("PostTemplate").ParseFiles("newpost.html"))
)

func main() {
	var blog1 = Blog{
		Name:        "Личный блог",
		Description: "Посты на разные темы",
		Posts: []Post{
			{ID: 1, Header: "Первый пост", Text: "13241 3g r1344 23 3fdgdg "},
			{ID: 2, Header: "Второй пост", Text: "hjghj 45456 435th hhth4 hth"},
			{ID: 3, Header: "Третий Пост", Text: "th4th4 t4ht 4th th 4tht ht"},
		},
	}
	router := http.NewServeMux()
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("src"))))
	router.HandleFunc("/blog", blog1.showBlog)
	router.HandleFunc("/blog/newpost", blog1.newPost)
	router.HandleFunc("/blog/", blog1.showPost)
	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func (blog1 *Blog) showBlog(w http.ResponseWriter, r *http.Request) {
	if err := tmplBlog.ExecuteTemplate(w, "blog", blog1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (blog1 *Blog) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	p, err := strconv.Atoi(path[2])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmplPost.ExecuteTemplate(w, "post", blog1.Posts[p-1]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (blog1 *Blog) newPost(w http.ResponseWriter, r *http.Request) {
	lenBlog := len(blog1.Posts)
	newpost := Post{
		ID:     lenBlog + 1,
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}

	if len(newpost.Header) != 0 {
		blog1.Posts = append(blog1.Posts, newpost)
	}

	if err := tmplNewPost.ExecuteTemplate(w, "newpost", blog1.Posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
