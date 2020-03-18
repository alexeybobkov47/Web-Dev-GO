package routers

import (
	"Web-Dev-GO/git/hw7/Beego_blog/controllers"
	"context"
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.39:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongo-db connected")
	if err = db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	// defer db.Disconnect(context.Background)

	beego.Router("/blog", &controllers.BlogController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/post/:id", &controllers.PostController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/edit/:id", &controllers.EditController{
		Controller: beego.Controller{},
		DB:         db,
	})
	beego.Router("/new", &controllers.NewPostController{
		Controller: beego.Controller{},
		DB:         db,
	})

}
