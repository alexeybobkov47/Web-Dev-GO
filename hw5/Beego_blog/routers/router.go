package routers

import (
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
	log.Println(db)
	defer db.Close()

	beego.Router("/blog", &controllers.blogController{
		Controller: beego.Controller{},
		Db:         db,
	})
	// beego.Router("/post/", &controllers.MainController{})
	// beego.Router("/editpost", &controllers.MainController{})
}
