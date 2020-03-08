package main

import (
	"database/sql"
	"fmt"
	"log"
)

func getBlogs(database *sql.DB) ([]Blog, error) {
	res := []Blog{}

	rows, err := database.Query("select * from Blog.Blogs")
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		blogs := Blog{}

		err := rows.Scan(&blogs.ID, &blogs.Name, &blogs.Description)
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, blogs)
	}

	return res, nil
}

func getPosts(database *sql.DB, id int) (Blog, error) {
	blogs := Blog{}

	row := database.QueryRow(fmt.Sprintf("select * from Blog.Blogs where Blogs.id = %v", id))
	err := row.Scan(&blogs.ID, &blogs.Name, &blogs.Description)
	if err != nil {
		log.Println(err)
		return blogs, err
	}

	rows, err := database.Query(fmt.Sprintf("select * from Blog.Posts WHERE Post.blog_id = %v", id))
	if err != nil {
		log.Println(err)
		return blogs, err
	}
	defer rows.Close()

	for rows.Next() {
		posts := Post{}

		err := rows.Scan(&posts.ID, new(int), &posts.Header, &posts.Text)
		if err != nil {
			log.Println(err)
			continue
		}

		blogs.Posts = append(blogs.Posts, posts)
	}

	return blogs, nil
}
