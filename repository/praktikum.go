package repository

import (
	"context"
	"fmt"
	"inibackend/config"
	"inibackend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByUsername(ctx context.Context, username string) (*model.UserLogin, error) {
	userCollection := config.MongoConnect(config.DBName).Collection(config.UserCollection)

	var user model.UserLogin
	filter := bson.M{"username": username}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("username %s tidak ditemukan", username)
		}
		return nil, err
	}

	return &user, nil
}

func InsertUser(ctx context.Context, user model.UserLogin) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UserCollection)

	// Cek apakah username sudah ada
	filter := bson.M{"username": user.Username}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("username %s sudah digunakan", user.Username)
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}