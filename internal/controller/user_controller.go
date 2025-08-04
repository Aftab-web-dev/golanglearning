package controller

import (
	"context"
	"log"
	"time"

	"github.com/Aftab-web-dev/learningproject/config"
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// CreateUserController handles the creation of a new user in the database
func CreateUserController(user models.User) (primitive.ObjectID, error) {
	
	// Check if the user already exists
	existingUser := models.User{}
	
	userCollection := config.DB.Collection("users")

	err := userCollection.FindOne(context.Background(), models.User{Email: user.Email}).Decode(&existingUser)
	if err == nil {	
		log.Fatal("User already exists")
		return primitive.NilObjectID, nil // User already exists, return nil ID	
	}

	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {	
		log.Fatal("Error creating user:", err)
		return primitive.NilObjectID, err // Return nil ID on error
	}
	return user.ID, nil

}