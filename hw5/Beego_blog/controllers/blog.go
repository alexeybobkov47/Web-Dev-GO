package controllers

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego"
)

// BlogController -
type BlogController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *BlogController) Get() {
	blogs, err := getBlogs(c.DB)
	if err != nil {
		log.Println(err)
		return
	}
	c.Data["Blog"] = blogs
	c.Data["Post"] = blogs.Posts
	c.TplName = "index.tpl"
}

// func (database *Server) newPost(w http.ResponseWriter, r *http.Request) {
// 	newpost := Post{
// 		Header: r.FormValue("header"),
// 		Text:   r.FormValue("text"),
// 	}

// 	if len(newpost.Header) != 0 && len(newpost.Text) != 0 {
// 		err := newPost(database.db, newpost)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}

// 	if err := tmplNewPost.ExecuteTemplate(w, "newpost", newpost); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
