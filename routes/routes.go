package routes

import (
    "github.com/gin-gonic/gin"
    "map/controllers"
)

func AuthRoutes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controllers.Signup)
        auth.POST("/login", controllers.Login) // Add the login endpoint
       auth.GET("/user/:id", controllers.GetUserProfile) // ID of the user

    }
}
