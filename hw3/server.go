package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("BlogTemplate").ParseFiles("index.html"))

var blog1 = Blog{
	Name:        "Личный блог",
	Description: "Посты на разные темы",
	Posts: []Post{
		{ID: "1", Header: "Первый пост", Text: "13241 3g r1344 23 3fdgdg "},
		{ID: "2", Header: "Второй пост", Text: "hjghj 45456 435th hhth4 hth"},
		{ID: "3", Header: "Третий Пост", Text: "th4th4 t4ht 4th th 4tht ht"},
	},
}

func main() {
	router := http.NewServeMux()
	port := "8080"
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("src"))))
	router.HandleFunc("/blog", showBlog)
	router.HandleFunc("/blog/1", showBlog)

	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func showBlog(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	if err := tmpl.ExecuteTemplate(w, "blog", blog1); err != nil {
		log.Println(err)
	}
}
