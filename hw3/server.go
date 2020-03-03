package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.New("BlogTemplate").ParseFiles("src/page.html"))

var blog1 = Blog{
	Name:        "Личный блог",
	Description: "Посты на разные темы",
	Posts: []Post{
		{ID: "1", Header: "Первый пост", Text: "13241 3g r1344 23 3fdgdg "},
		{ID: "2", Header: "Вторая пост", Text: "hjghj 45456 435th hhth4 hth"},
		{ID: "3", Header: "Третий Пост", Text: "th4th4 t4ht 4th th 4tht ht"},
	},
}

func main() {
	router := http.NewServeMux()
	port := "8080"
	router.HandleFunc("/", showBlog)
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func showBlog(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	}
	w.Header().Add("Content-Type", contentType)

	if err := tmpl.ExecuteTemplate(w, "blog", blog1); err != nil {
		log.Println(err)
	}
}
