package controllers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/models"
	"database/sql"
	"log"

	"github.com/astaxie/beego"
)

// NewPostController -
type NewPostController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *NewPostController) Get() {
	c.TplName = "newpost.tpl"
}

// Post -
func (c *NewPostController) Post() {
	newpost := models.Post{
		Header: c.Ctx.Request.FormValue("header"),
		Text:   c.Ctx.Request.FormValue("text"),
	}
	if len(newpost.Header) != 0 && len(newpost.Text) != 0 {
		err := newPost(c.DB, newpost)
		if err != nil {
			log.Println(err)
			return
		}
	}
	c.TplName = "newpost.tpl"
	c.Redirect("/blog", 301)
}
