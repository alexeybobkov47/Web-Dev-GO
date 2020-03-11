package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	blogs, err := getBlogs(database.db)
	if err != nil {
		log.Println(err)
		return
	}

	c.Data["Blog"] = blogs
	c.TplName = "index.tpl"
}
