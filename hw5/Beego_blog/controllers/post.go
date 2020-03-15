package controllers

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego"
)

// PostController -
type PostController struct {
	beego.Controller
	DB *sql.DB
}

// Get -
func (c *PostController) Get() {
	post, err := getPost(c.DB, c.Ctx.Input.Param(":id"))
	if err != nil {
		log.Println(err)
		c.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Data["Post"] = post
	c.TplName = "post.tpl"

}
