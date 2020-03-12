package routers

import (
	"Web-Dev-GO/git/hw5/Beego_blog/controllers"
	"database/sql"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// const dsn = "root:12345@tcp(10.111.100.232:3306)/Site?charset=utf8"
const dsn = "root:12345@tcp(192.168.0.39:3306)/Site?charset=utf8"

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Router("/blog", &controllers.BlogController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/post/:id([0-9])+", &controllers.PostController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/edit/:id([0-9])+", &controllers.EditController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/new", &controllers.NewPostController{
		Controller: beego.Controller{},
		DB:         db,
	})

}
