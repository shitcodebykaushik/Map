package main

import (
    "context"
    "log"
    "map/db"
    "map/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize MongoDB connection
    db.Init()
    defer func() {
        if err := db.Client.Disconnect(context.Background()); err != nil {
            log.Fatal("Error disconnecting from MongoDB:", err)
        }
    }()

    // Set up Gin router
    router := gin.Default()

    // Register routes
    routes.AuthRoutes(router)

    // Start server on port 8080
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
