package controllers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/astaxie/beego"
)

// PostController -
type PostController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *PostController) Get() {
	path := strings.Split(c.Ctx.Request.URL.Path, "/")
	post, err := getPost(c.DB, (path[len(path)-1]))
	if err != nil {
		log.Println(err)
		return
	}
	c.Data["Post"] = post
	c.TplName = "post.tpl"

}
