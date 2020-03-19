package main

import (
	"database/sql"
	"log"
	"testing"
)

func TestNewPost(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	s := Server{
		db: db,
	}
	post := Post{
		Header: "Тестовый лист",
		Text:   "Описание тестового листа",
	}
	if err := newPost(s.db, post); err != nil {
		t.Error(err)
	}
}

func TestGetPost(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	s := Server{
		db: db,
	}
	type iD struct {
		id string
	}
	postID := iD{}
	row := s.db.QueryRow("select id from Site.Blogs WHERE Site.post.header = 'Тестовый лист'")
	err = row.Scan(&postID.id)
	log.Println(postID)

	post := Post{
		Header: "Тестовый лист",
		Text:   "Описание тестового листа",
	}
	receivedpost, err := getPosts(s.db, postID.id)
	if err != nil {
		t.Error(err)
	}
	if post != receivedpost {
		t.Error("Посты не совпадают")
	}

}
