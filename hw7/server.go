package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tmplBlog     = template.Must(template.New("BlogTemplate").ParseFiles("index.html"))
	tmplPost     = template.Must(template.New("PostTemplate").ParseFiles("post.html"))
	tmplNewPost  = template.Must(template.New("PostTemplate").ParseFiles("newpost.html"))
	tmplEditPost = template.Must(template.New("PostTemplate").ParseFiles("editpost.html"))
	// dsn          = "root:12345@tcp(192.168.0.39:3306)/Site?charset=utf8"
	dsn = "root:12345@tcp(10.111.100.232:3306)/Site?charset=utf8"
)

func main() {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	s := Server{
		db: db,
	}

	router := http.NewServeMux()
	router.HandleFunc("/blog", s.showBlog)
	router.HandleFunc("/blog/newpost", s.newPost)
	router.HandleFunc("/blog/", s.showPost)
	router.HandleFunc("/blog/edit/", s.editPost)
	port := "8080"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
