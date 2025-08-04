package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aftab-web-dev/learningproject/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	//Load environment variables
	err := godotenv.Load()
	 if err != nil {
        log.Fatal("Error loading .env file")
    }

	port := os.Getenv("PORT")
	if port == "" { 
		port = "8080"
	}

	// initialize mongoDB connection here
	mongoURI := os.Getenv("MONGO_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	
	// create a new MongoDB client
	client , err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	
	defer func() { 
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	} ()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	
	// Initialize Gin router
	r := gin.Default()
	routes.RegisterRoutes(r)
	
	// Custom 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":   "Route not found",
			"message": "The route you are trying to access does not exist",
		})
	})

	r.Run(":" + port)

	// Create HTTP server with Gin handler
	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	//Graceful shutdown setip
	done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGALRM)
    
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
    
	<-done

	slog.Info("Received shutdown signal, shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down server", "error", err.Error())
	} else {
		slog.Info("Server gracefully stopped")
	}


}
