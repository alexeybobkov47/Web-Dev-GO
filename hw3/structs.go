package main

// Blog - Структура блога
type Blog struct {
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
