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
		return blogs, err
	}

	rows, err := database.Query(fmt.Sprintf("select * from Site.Post"))
	if err != nil {
		return blogs, err
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

	return blogs, nil
}

func getPosts(database *sql.DB, id string) (Post, error) {
	posts := Post{}
	row := database.QueryRow(`select * from Site.Post WHERE Site.Post.id =` + id)
	err := row.Scan(&posts.ID, &posts.Header, &posts.Text)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func newPost(database *sql.DB, newpost Post) error {
	// log.Printf("insert into Site.Post (header, text) values ('%s','%s');", newpost.Header, newpost.Text)
	post := fmt.Sprintf("insert into Site.Post (header, text) values ('%s','%s');", newpost.Header, newpost.Text)
	// database.QueryRow("insert into Site.Post (header, text) values (" + newpost.Header + "," + newpost.Text + ")")
	_, err := database.Exec(post)
	return err

}

func editPost(database *sql.DB, editPost Post, id string) error {
	post := fmt.Sprintf("update Site.Post set header='%s', text='%s' where id=%s", editPost.Header, editPost.Text, id)
	_, err := database.Exec(post)
	return err
}

func deletePost(database *sql.DB, id string) error {
	deletepost := fmt.Sprintf(`delete from Site.Post WHERE Site.Post.id =` + id)
	_, err := database.Exec(deletepost)
	return err
}
