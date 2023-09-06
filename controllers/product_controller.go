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

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validateProduct = validator.New()

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var product models.Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: err.Error()})
			return
		}

		if validationErr := validateProduct.Struct(&product); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Validation Error", Data: validationErr.Error()})
			return
		}

		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		result, err := productCollection.InsertOne(ctx, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.ProductResponse{Status: http.StatusCreated, Message: "Product Created Successfully", Data: result.InsertedID})
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productID := c.Param("productId")
		var product models.Product
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: err.Error()})
			return
		}

		err = productCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "Product Not Found", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "Success", Data: product})
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productID := c.Param("productId")
		var updatedProduct models.Product
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: err.Error()})
			return
		}

		if err := c.BindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: err.Error()})
			return
		}

		if validationErr := validateProduct.Struct(&updatedProduct); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Validation Error", Data: validationErr.Error()})
			return
		}

		update := bson.M{"name": updatedProduct.Name, "price": updatedProduct.Price, "stock": updatedProduct.Stock, "image": updatedProduct.Image, "updatedAt": time.Now()}
		result, err := productCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: err.Error()})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "Product Not Found", Data: "No documents were modified."})
			return
		}

		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "Product Updated Successfully", Data: updatedProduct})
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		productID := c.Param("productId")
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: err.Error()})
			return
		}

		result, err := productCollection.DeleteOne(ctx, bson.M{"_id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: err.Error()})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "Product Not Found", Data: "No documents were deleted."})
			return
		}

		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "Product Deleted Successfully", Data: nil})
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var products []models.Product
		defer cancel()

		cursor, err := productCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: err.Error()})
			return
		}

		if err = cursor.All(ctx, &products); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "Success", Data: products})
	}
}
