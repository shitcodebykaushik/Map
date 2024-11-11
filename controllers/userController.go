// controllers/userController.go
package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "map/db"
    "map/models"
)

func GetUserProfile(c *gin.Context) {
    userID := c.Param("id") // Use user ID from request URL or token
    var user models.User

    userCollection := db.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "full_name": user.FullName,
        "email":     user.Email,
         // Assuming you have an avatar field
    })
}
