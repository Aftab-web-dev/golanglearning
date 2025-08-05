package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aftab-web-dev/learningproject/config"
	"github.com/Aftab-web-dev/learningproject/internal/middleware"
	"github.com/Aftab-web-dev/learningproject/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }

    // Connect to MongoDB
    config.ConnectMongoDB()
    defer config.DisconnectMongoDB() // ✅ Disconnect only on shutdown

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Initialize Gin router
    r := gin.Default()
    r.Use(middleware.ContextTimeoutMiddleware(5 * time.Second))
    routes.UserRoutes(r)
    // Custom 404 handler
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "error":   "Route not found",
            "message": "The route you are trying to access does not exist",
        })
    })

    // Create HTTP server with Gin handler
    server := &http.Server{
        Addr:    ":" + port, // ✅ Fixed
        Handler: r,
    }

    // Channel to listen for OS signals
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

    // Start server in a goroutine
    go func() {
        slog.Info("Server started on port " + port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

    // Wait for termination signal
    <-quit
    slog.Info("Received shutdown signal, shutting down server...")

    // Context for graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        slog.Error("Error shutting down server", "error", err.Error())
    } else {
        slog.Info("Server gracefully stopped")
    }
}
