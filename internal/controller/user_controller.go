package controller

import (
	"context"
	"fmt"

	"github.com/Aftab-web-dev/learningproject/config"
	"github.com/Aftab-web-dev/learningproject/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserController(ctx context.Context, user models.User) (bson.ObjectID, error) {
	//Password Hashing
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return bson.ObjectID{}, err
	}

	userCollection := config.DB.Collection("users")

	user.ID = bson.NewObjectID()
	user.Password = string(hashPassword)

	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
	if err == nil {
		return bson.NilObjectID, fmt.Errorf("email already exists")
	} else if err != mongo.ErrNoDocuments {
		return bson.NilObjectID, err // some other DB error
	}

	// Check username uniqueness
	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)
	if err == nil {
		return bson.NilObjectID, fmt.Errorf("username already taken")
	} else if err != mongo.ErrNoDocuments {
		return bson.NilObjectID, err
	}

	//  Insert new user
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("InsertOne error: %v\n", err)
		return bson.NilObjectID, err
	}

	return user.ID, nil
}

func GetallUsersController(ctx context.Context) ([]models.User, error) {
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

func GetUserbyidController(ctx context.Context, id string) (models.User, error) {
	userCollection := config.DB.Collection("users")
	var foundUser models.User

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return foundUser, fmt.Errorf("invalid user ID format")
	}
	err = userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&foundUser)
	if err != nil {
		return foundUser, fmt.Errorf("user not found")
	}
	return foundUser, nil
}

func LoginuserController(ctx context.Context, user models.LoginUser) error {
	userCollection := config.DB.Collection("users")

	var dbUser models.User

	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&dbUser)

	if err != nil {
		return fmt.Errorf("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func DeleteUserbyIdController(ctx context.Context, id string) error {
	userCollection := config.DB.Collection("users")

	ObjID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format")
	}

	//Perform delete opertaion
	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": ObjID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func UpdatedetailsController(ctx context.Context, id string, update models.UserUpdate) error {
    userCollection := config.DB.Collection("users")

    objID, err := bson.ObjectIDFromHex(id)
    if err != nil {
        return fmt.Errorf("invalid user ID format")
    }

    updateFields := bson.M{}

    if update.Username != nil {
        updateFields["username"] = *update.Username
    }
    if update.Email != nil {
        updateFields["email"] = *update.Email
    }
    if update.Phonenumber != nil {
        updateFields["phone_number"] = *update.Phonenumber
    }
    if update.Password != nil {
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
        updateFields["password"] = string(hashedPassword)
    }

    if len(updateFields) == 0 {
        return fmt.Errorf("no fields to update")
    }

    result, err := userCollection.UpdateByID(ctx, objID, bson.M{"$set": updateFields})
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return fmt.Errorf("user not found")
    }

    return nil
}
