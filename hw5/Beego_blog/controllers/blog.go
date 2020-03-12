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
