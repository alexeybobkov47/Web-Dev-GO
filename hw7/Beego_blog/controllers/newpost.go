package controllers

import (
	"Web-Dev-GO/git/hw7/Beego_blog/models"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewPostController -
type NewPostController struct {
	beego.Controller
	DB *mongo.Client
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
