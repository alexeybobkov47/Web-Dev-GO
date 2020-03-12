package controllers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/models"
	"database/sql"
	"log"
	"strings"

	"github.com/astaxie/beego"
)

// EditController -
type EditController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *EditController) Get() {
	path := strings.Split(c.Ctx.Request.URL.Path, "/")
	p := (path[len(path)-1])
	posts, err := getPost(c.DB, p)
	if err != nil {
		log.Println(err)
		return
	}
	editpost := models.Post{
		Header: c.Ctx.Request.FormValue("header"),
		Text:   c.Ctx.Request.FormValue("text"),
	}
	if len(editpost.Header) != 0 && len(editpost.Text) != 0 {
		err := editPost(c.DB, editpost, p)
		if err != nil {
			log.Println(err)
			return
		}
	}

	c.Data["Post"] = posts
	c.TplName = "editpost.tpl"
}

// Post -
func (c *EditController) Post() {
	path := strings.Split(c.Ctx.Request.URL.Path, "/")
	err := deletePost(c.DB, (path[len(path)-1]))
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect("/blog", 301)
}

// Delete -
// func (c *EditController) Delete() {
// 	id := strings.Join((c.Ctx.Request.URL.Query())["id"], "")
// 	err := deletePost(c.DB, id)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	c.Redirect("/Blog", 301)
// }
