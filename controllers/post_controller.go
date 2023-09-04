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
			Id:        primitive.NewObjectID(),
			Title:     post.Title,
			Content:   post.Content,
			Image:     post.Image,
			UserId:    post.UserId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Views:     1,
			Pongs:     1,
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

		updateResult, updateErr := postCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$inc": bson.M{"views": 1}})
		if updateErr != nil {

			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": updateErr.Error()}})
			return
		}

		if updateResult.ModifiedCount == 0 {

			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "No documents were modified."}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": post}})
	}
}

func IncreaseViews() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(postId)

		result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$inc": bson.M{"views": 1000}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
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

func GetPostsByViews() gin.HandlerFunc {
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

func GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		postId := c.Param("postId")
		objId, err := primitive.ObjectIDFromHex(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var post models.Post
		err = postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": post}})
	}
}

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		postID := c.Param("postId")
		objID, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var comment models.Comment

		if err := c.BindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		comment.Id = primitive.NewObjectID()
		comment.CreatedAt = time.Now()
		comment.UpdatedAt = time.Now()
		comment.Pongs = 1

		updateResult, updateErr := postCollection.UpdateOne(
			ctx,
			bson.M{"_id": objID},
			bson.M{"$push": bson.M{"comments": comment}},
		)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"data": updateErr.Error()}})
			return
		}

		if updateResult.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, responses.PostResponse{Status: http.StatusNotFound, Message: "Post Not Found", Data: map[string]interface{}{"data": "The post with the specified ID was not found."}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "Comment Created Successfully", Data: map[string]interface{}{"data": comment}})
	}
}

func EditComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		postID := c.Param("postId")
		commentID := c.Param("commentId")
		objPostID, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedComment models.Comment

		if err := c.BindJSON(&updatedComment); err != nil {
			c.JSON(http.StatusBadRequest, responses.PostResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		updateResult, updateErr := postCollection.UpdateOne(
			ctx,
			bson.M{"_id": objPostID, "comments._id": commentID},
			bson.M{"$set": bson.M{"comments.$.content": updatedComment.Content}},
		)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, responses.PostResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"data": updateErr.Error()}})
			return
		}

		if updateResult.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, responses.PostResponse{Status: http.StatusNotFound, Message: "Comment Not Found", Data: map[string]interface{}{"data": "The comment with the specified ID was not found in the post."}})
			return
		}

		c.JSON(http.StatusOK, responses.PostResponse{Status: http.StatusOK, Message: "Comment Updated Successfully", Data: map[string]interface{}{"data": updatedComment}})
	}
}
