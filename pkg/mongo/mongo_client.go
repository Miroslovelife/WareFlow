package mongo

import (
	"context"
	"log"
	"time"

	"github.com/Miroslovelife/WareFlow/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI    config.MongoURI
	DBName config.DBName
}

type MongoClient struct {
	client *mongo.Client
	dbName string
}

// NewMongoClient создает новый экземпляр MongoClient
func NewMongoClient(cfg MongoConfig) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(string(cfg.URI)))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return &MongoClient{
		client: client,
		dbName: string(cfg.DBName),
	}, nil
}

// NewMongoCollection возвращает коллекцию с указанным именем
func NewMongoCollection(client *MongoClient, collectionName string) *mongo.Collection {
	return client.client.Database(client.dbName).Collection(collectionName)
}

// Close закрывает соединение с MongoDB
func (m *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.client.Disconnect(ctx)
}
