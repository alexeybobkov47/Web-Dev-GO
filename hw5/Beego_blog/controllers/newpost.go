package controllers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/models"
	"database/sql"

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
			c.Ctx.ResponseWriter.WriteHeader(500)
			_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
			return
		}
	}
	c.TplName = "newpost.tpl"
	c.Redirect("/blog", 301)
}
