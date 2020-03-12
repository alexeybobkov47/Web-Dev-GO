package models

import "database/sql"

// Blog - Структура блога
type Blog struct {
	ID          int
	Name        string
	Description string
	Posts       []Post
}

//Post - Структура постов в блоге
type Post struct {
	ID     int
	Header string
	Text   string
}

// Server -
type Server struct {
	db *sql.DB
}
