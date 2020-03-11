package routers

import (
	"Beego_blog/controllers"
	"database/sql"
	"log"

	"github.com/astaxie/beego"
)

const dsn = "root:12345@tcp(10.111.100.232:3306)/Site?charset=utf8"

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	beego.Router("/blog", &controllers.MainController{})
	beego.Router("/post/", &controllers.MainController{})
	beego.Router("/editpost", &controllers.MainController{})
}
