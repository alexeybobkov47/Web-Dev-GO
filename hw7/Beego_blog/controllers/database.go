package controllers

import (
	"Web-Dev-GO/git/hw7/Beego_blog/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getBlogs(database *mongo.Client) (models.Blog, error) {
	c := database.Database("myblog").Collection("blog")
	data, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return models.Blog{}, err
	}
	blogs := []models.Blog{}
	if err := data.All(context.Background(), &blogs); err != nil {
		return models.Blog{}, err
	}
	blog := blogs[0]
	p := database.Database("myblog").Collection("posts")
	rows, err := p.Find(context.Background(), bson.D{})
	if err != nil {
		return models.Blog{}, err
	}
	if err := rows.All(context.Background(), &blog.Posts); err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func getPost(database *mongo.Client, id string) (models.Post, error) {
	log.Println(id)
	p := database.Database("myblog").Collection("posts")
	filter := bson.D{{Key: "header", Value: id}}
	rows := p.FindOne(context.Background(), filter)
	post := new(models.Post)
	if err := rows.Decode(post); err != nil {
		return models.Post{}, err
	}
	return *post, nil
}

func newPost(database *mongo.Client, newpost models.Post) error {
	c := database.Database("myblog").Collection("posts")
	_, err := c.InsertOne(context.Background(), newpost)
	return err

}

func editPost(database *mongo.Client, editPost models.Post, id string) error {
	filter := bson.D{{Key: "header", Value: id}}
	update := bson.D{}
	if len(editPost.Header) != 0 {
		update = append(update, bson.E{Key: "header", Value: editPost.Header})
	}
	if len(editPost.Text) != 0 {
		update = append(update, bson.E{Key: "text", Value: editPost.Text})
	}
	update = bson.D{{Key: "$set", Value: update}}
	c := database.Database("myblog").Collection("posts")
	_, err := c.UpdateOne(context.Background(), filter, update)

	return err

}

// DeletePost - удаление поста
func deletePost(database *mongo.Client, id string) error {
	c := database.Database("myblog").Collection("posts")
	filter := bson.D{{Key: "header", Value: id}}
	_, err := c.DeleteOne(context.Background(), filter)
	return err
}
