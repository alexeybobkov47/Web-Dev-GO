package controllers

import (
	"database/sql"
	"log"
	"strings"

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

// PostController -
type PostController struct {
	beego.Controller
	DB *sql.DB
}

func (c *PostController) showPost() {
	path := strings.Split(c.Ctx.Request.URL.Path, "/")
	post, err := getPost(c.DB, (path[len(path)-1]))
	if err != nil {
		log.Println(err)
		return
	}
	c.Data["Post"] = post
	c.TplName = "post.tpl"

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

// func (database *Server) editPost(w http.ResponseWriter, r *http.Request) {
// 	path := strings.Split(r.URL.Path, "/")
// 	p := (path[len(path)-1])
// 	posts, err := getPosts(database.db, p)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	delPost := r.URL.Query()
// 	if delPost["delete"] != nil {
// 		err := deletePost(database.db, strings.Join(delPost["id"], ""))
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}

// 	editpost := Post{
// 		Header: r.FormValue("header"),
// 		Text:   r.FormValue("text"),
// 	}
// 	if len(editpost.Header) != 0 && len(editpost.Text) != 0 {
// 		err := editPost(database.db, editpost, p)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}

// 	if err := tmplEditPost.ExecuteTemplate(w, "editpost", posts); err != nil {
// 		log.Println(err)
// 		return
// 	}
// }
