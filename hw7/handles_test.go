package main

import (
	"database/sql"
	"log"
	"net/http/httptest"
	"testing"
)

func TestSHTTP(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
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
