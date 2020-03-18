package controllers

import (
	"Web-Dev-GO/git/hw7/Beego_blog/models"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func testNewPost(t *testing.T) {

	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.39:27017"))
	if err != nil {
		log.Fatal(err)
	}

	post := models.Post{
		Header: "Тестовый заголовок",
		Text:   "Текст тестовый",
	}

	if err := newPost(db, post); err != nil {
		t.Error(err)
	}
}
