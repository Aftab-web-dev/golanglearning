package config

import (
    "context"
    "fmt"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database
var MongoClient *mongo.Client

func ConnectMongoDB() {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb+srv://username:password@cluster.mongodb.net"
    }

    clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
    client, err := mongo.Connect(clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping to confirm connection
    if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    // Store global client and DB
    MongoClient = client
    DB = client.Database("testdb") // Change to your database name
}

// Call this in main.go during shutdown
func DisconnectMongoDB() {
    if MongoClient != nil {
        if err := MongoClient.Disconnect(context.Background()); err != nil {
            log.Println("Error disconnecting MongoDB:", err)
        }
    }
}
