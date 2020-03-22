package main

import (
	"database/sql"
	"log"
	"net/http/httptest"
	"testing"
)

func TestShowBlog(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	s := Server{
		db: db,
	}

	url := "http://localhost:8080/blog"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	s.showBlog(w, req)
	if w.Code != 200 {
		t.Errorf("wrong StatusCode: got %d, expected 200",
			w.Code)
	}
	log.Printf("StatusCode: got %d", w.Code)
}
func TestShowPost(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	s := Server{
		db: db,
	}

	url := "http://localhost:8080/post/20"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	s.showPost(w, req)
	if w.Code != 200 {
		t.Errorf("wrong StatusCode: got %d, expected 200",
			w.Code)
	}
	log.Printf("StatusCode: got %d", w.Code)
}
