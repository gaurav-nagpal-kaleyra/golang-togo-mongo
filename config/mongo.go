package config

import (
	"context"
	"fmt"
	"golang-todo/constants"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Initialize mongodb connection here.

var MongoConn *mongo.Client

func ConnectMongo() error {

	connectURI := fmt.Sprintf("mongodb://%s:%s", os.Getenv(constants.MongoDbHost), os.Getenv(constants.MongoDbPort))

	options.Client().ApplyURI(connectURI)

	var err error
	MongoConn, err = mongo.Connect(context.Background(), options.Client())
	if err != nil {
		fmt.Println("Error connecting to mongodb", err.Error())
		return err
	}

	fmt.Println("MongoDB connected successfully")
	return nil
}
