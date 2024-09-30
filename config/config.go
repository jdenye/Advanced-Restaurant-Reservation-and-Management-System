package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

var DB *mongo.Client

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client
    fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Database(os.Getenv("DB_NAME")).Collection(collectionName)
}
