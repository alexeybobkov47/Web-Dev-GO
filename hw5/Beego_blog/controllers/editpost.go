package controllers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/models"
	"database/sql"

	"github.com/astaxie/beego"
)

// EditController -
type EditController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *EditController) Get() {
	p := c.Ctx.Input.Param(":id")
	posts, err := getPost(c.DB, p)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	editpost := models.Post{
		Header: c.Ctx.Request.FormValue("header"),
		Text:   c.Ctx.Request.FormValue("text"),
	}
	if len(editpost.Header) != 0 && len(editpost.Text) != 0 {
		err := editPost(c.DB, editpost, p)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(500)
			_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
			return
		}
	}

	c.Data["Post"] = posts
	c.TplName = "editpost.tpl"
}

// Post -
func (c *EditController) Post() {

	err := deletePost(c.DB, c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Redirect("/blog", 301)
}
