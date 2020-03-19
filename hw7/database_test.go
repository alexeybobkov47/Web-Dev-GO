package main

import (
	"database/sql"
	"log"
	"testing"
)

func getDbCon() (s Server) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	s := Server{
		db: db,
	}
	return s
}

func TestNewPost(t *testing.T) {

	post := Post{
		ID:     9999999,
		Header: "Тестовый лист",
		Text:   "Описание тестового листа",
	}
	db := getDbCon()
	if err := newPost(db, post); err != nil {
		t.Error(err)
	}
}

func TestGetPost(t *testing.T) {
	id := "9999999"
	post := Post{
		ID:     9999999,
		Header: "Тестовый лист",
		Text:   "Описание тестового листа",
	}
	db := getDbCon()
	receivedpost, err := getPosts(db, id)
	if err != nil {
		t.Error(err)
	}
	if post != receivedpost {
		t.Error("Посты не совпадают")
	}

}
