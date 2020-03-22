package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostController -
type PostController struct {
	beego.Controller
	DB *mongo.Client
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
