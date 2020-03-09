package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getBlogs(database *sql.DB) (Blog, error) {
	blogs := Blog{}
	row := database.QueryRow("select * from Site.Blogs")
	err := row.Scan(&blogs.ID, &blogs.Name, &blogs.Description)
	if err != nil {
		log.Println(err)
	}

	rows, err := database.Query(fmt.Sprintf("select * from Site.Post"))
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		posts := Post{}

		err := rows.Scan(&posts.ID, &posts.Header, &posts.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		blogs.Posts = append(blogs.Posts, posts)
	}

	return blogs, err
}

func getPosts(database *sql.DB, id string) (Post, error) {
	posts := Post{}
	log.Println("select * from Site.Post WHERE Site.Post.id =" + id)
	row := database.QueryRow("select * from Site.Post WHERE Site.Post.id =" + id)
	err := row.Scan(&posts.ID, &posts.Header, &posts.Text)
	if err != nil {
		log.Println(err)
	}
	return posts, err
}
