package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

// BlogController -
type BlogController struct {
	beego.Controller
	DB *mongo.Client
}

// Get -
func (c *BlogController) Get() {
	blogs, err := getBlogs(c.DB)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Data["Blog"] = blogs
	c.Data["Post"] = blogs.Posts
	c.TplName = "index.tpl"
}
