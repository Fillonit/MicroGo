package controllers

import (
	"MicroGo/configs"
	"MicroGo/models"
	"MicroGo/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
var validatePost = validator.New()

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var post models.Post
		defer cancel()

		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validatePost.Struct(&post); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newPost := models.Post{
			Id:      primitive.NewObjectID(),
			Title:   post.Title,
			Content: post.Content,
			Image:   post.Image,
			UserId:  post.UserId,
		}

		result, err := postCollection.InsertOne(ctx, newPost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PostResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		var post models.Post
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(postId)

		err := postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": post}})
	}
}

func GetAllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var posts []models.Post
		defer cancel()

		cursor, err := postCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if err = cursor.All(ctx, &posts); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": posts}})
	}
}

func EditAPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		var post models.Post
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(postId)

		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validatePost.Struct(&post); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"title": post.Title, "content": post.Content, "image": post.Image, "userId": post.UserId}
		result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedPost models.Post
		if result.MatchedCount == 1 {
			err := postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPost}})
	}
}

func DeleteAPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(postId)

		result, err := postCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
