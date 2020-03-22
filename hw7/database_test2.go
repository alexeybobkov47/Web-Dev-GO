package main

import (
	"database/sql"
	"testing"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	s := Server{
		db: db,
	}
	post := Post{
		Header: "Тестовый заголовок",
		Text:   "Тест",
	}

	updatepost := Post{
		Header: "Обновленный заголовок",
		Text:   "Обновленный тест",
	}

}
