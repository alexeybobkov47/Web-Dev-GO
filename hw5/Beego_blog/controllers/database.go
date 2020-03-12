package controllers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/models"
	"database/sql"
	"fmt"
	"log"
)

func getBlogs(database *sql.DB) (models.Blog, error) {
	blogs := models.Blog{}
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
		posts := models.Post{}
		err := rows.Scan(&posts.ID, &posts.Header, &posts.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		blogs.Posts = append(blogs.Posts, posts)
	}

	return blogs, nil
}

func getPost(database *sql.DB, id string) (models.Post, error) {
	post := models.Post{}
	row := database.QueryRow(`select * from Site.Post WHERE Site.Post.id =` + id)
	err := row.Scan(&post.ID, &post.Header, &post.Text)
	if err != nil {
		return post, err
	}
	return post, nil
}

func newPost(database *sql.DB, newpost models.Post) error {
	post := fmt.Sprintf("insert into Site.Post (header, text) values ('%s','%s');", newpost.Header, newpost.Text)
	_, err := database.Exec(post)
	return err

}

func editPost(database *sql.DB, editPost models.Post, id string) error {
	post := fmt.Sprintf("update Site.Post set header='%s', text='%s' where id=%s", editPost.Header, editPost.Text, id)
	_, err := database.Exec(post)
	return err
}

// DeletePost - удаление поста
func deletePost(database *sql.DB, id string) error {
	deletepost := fmt.Sprintf(`delete from Site.Post WHERE Site.Post.id =` + id)
	_, err := database.Exec(deletepost)
	return err
}
