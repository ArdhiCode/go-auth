package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ArdhiCode/go-auth/internal/config"
	"github.com/ArdhiCode/go-auth/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := config.DB.Collection("users").CountDocuments(ctx, bson.M{"email": user.Email})

	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Email already exist.")
	}

	_, err = config.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
