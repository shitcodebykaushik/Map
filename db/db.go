package db

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the MongoDB client instance
var Client *mongo.Client

func Init() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Set client options with URI from environment variable
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

    // Connect to MongoDB
    var err error
    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
    defer cancel()
    if err = Client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
    return Client.Database("map").Collection(collectionName)
}
