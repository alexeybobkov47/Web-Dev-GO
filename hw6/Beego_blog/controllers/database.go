package controllers

import (
	"Web-Dev-GO/git/hw6/Beego_blog/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// > db.blog.insert({"name": "Личный блог", "description": "Описание", posts: [{"header": "1111", "text": "11111"}, {"header": "2222", "text": "22222"}]})
func getBlogs(database *mongo.Client) (models.Blog, error) {
	c := database.Database("myblog").Collection("blog")
	data, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return models.Blog{}, err
	}
	blogs := make([]models.Blog, 0, 1)
	if err := data.All(context.Background(), &blogs); err != nil {
		return models.Blog{}, err
	}
	return blogs[0], nil
}

func getPost(database *mongo.Client, id string) (models.Post, error) {
	p := database.Database("myblog").Collection("blog")
	// filter := bson.M{"posts": bson.M{"$elemMatch": bson.M{"header": "1111"}}}
	filter := bson.D{{Key: "posts.header", Value: id}}
	// rows := p.FindOne(context.Background(), filter)
	rows := p.FindOne(context.Background(), filter)
	// blog := new(models.Blog)
	post := models.Post{}
	if err := rows.Decode(&post); err != nil {
		return models.Post{}, err
	}
	log.Println(post)
	return post, nil
}

func newPost(database *mongo.Client, newpost models.Post) error {
	c := database.Database("myblog").Collection("blog")
	_, err := c.InsertOne(context.Background(), bson.D{{"posts": bson.D{"header": newpost.Header, "text": newpost.Text}}})
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
