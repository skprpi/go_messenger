package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoSettings struct {
	uri               string
	databaseName      string
	collectionName    string
	connectionTimeout time.Duration
}

func NewMongoSettings(mongoUri string, databaseName string, collectionName string, connectionTimeout time.Duration) *MongoSettings {
	return &MongoSettings{
		uri:               mongoUri,
		databaseName:      databaseName,
		collectionName:    collectionName,
		connectionTimeout: connectionTimeout,
	}
}

func NewMongoCollection(ctx context.Context, settings *MongoSettings) (*mongo.Collection, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, settings.connectionTimeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection := client.Database(settings.databaseName).Collection(settings.collectionName)
	return collection, cancel
}