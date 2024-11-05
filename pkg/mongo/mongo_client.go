package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	Client *mongo.Client
	DBName string
}

func NewMongoClient(client *mongo.Client, dbName string) *MongoClient {
	return &MongoClient{
		Client: client,
		DBName: dbName,
	}
}

func (m *MongoClient) Collection(collectionName string) *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(collectionName)
}
