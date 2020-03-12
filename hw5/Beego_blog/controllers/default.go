package controllers

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego"
)

type blogController struct {
	beego.Controller
	db *sql.DB
}

func (c *blogController) Get() {

	blogs, err := getBlogs(c.db)
	if err != nil {
		log.Println(err)
		return
	}

	c.Data["Blog"] = blogs
	c.TplName = "index.tpl"
}
