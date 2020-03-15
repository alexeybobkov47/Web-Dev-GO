package controllers

import (
	"database/sql"

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
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Data["Blog"] = blogs
	c.Data["Post"] = blogs.Posts
	c.TplName = "index.tpl"
}
