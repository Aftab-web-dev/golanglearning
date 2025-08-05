package controller

import (
	"context"
	"fmt"

	"github.com/Aftab-web-dev/learningproject/config"
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateUserController(ctx  context.Context, user models.User) (primitive.ObjectID, error) {
	userCollection := config.DB.Collection("users")
	user.ID = primitive.NewObjectID()

	//  Check email uniqueness
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return primitive.NilObjectID, fmt.Errorf("email already exists")
	} else if err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, err // some other DB error
	}

	// Check username uniqueness
	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return primitive.NilObjectID, fmt.Errorf("username already taken")
	} else if err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, err
	}

	//  Insert new user
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("InsertOne error: %v\n", err)
		return primitive.NilObjectID, err
	}

	return user.ID, nil
}

func GetallUsersController(ctx  context.Context) ([]models.User, error) {
	userCollection := config.DB.Collection("users")
	cursor, err := userCollection.Find(ctx, bson.M{}) // Find all users

	if err != nil {
		return nil, err 
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	 if users == nil {
        users = []models.User{}
    }
   return users, nil
} 

func GetUserbyidController(ctx  context.Context, id string ) (models.User, error) {
	userCollection := config.DB.Collection("users")
	var foundUser models.User

	objID , err := primitive.ObjectIDFromHex(id)
	if err != nil { 
		return foundUser, fmt.Errorf("invalid user ID format")
	}
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&foundUser)
	if err != nil { 
		return foundUser, fmt.Errorf("user not found")
	}
	return foundUser, nil
}
