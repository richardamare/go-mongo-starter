package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var mongoClient *mongo.Client
var dbName string

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(name)
}

func ExistsByID(collection, id string) error {
	bId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bId}
	count, err := GetCollection(collection).CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("not found")
	}
	return nil
}

func ExistsByField(collection, field, value string) error {
	filter := bson.M{field: value}
	count, err := GetCollection(collection).CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("not found")
	}
	return nil
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return errors.New("you must set your 'DATABASE' environmental variable")
	} else {
		dbName = database
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return nil
}

func CloseMongoDB() {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
