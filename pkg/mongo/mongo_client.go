package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoClient создает новый экземпляр MongoClient
func NewMongoClient(uri, databaseName string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Подключаемся к MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Проверяем соединение
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	// Выбираем базу данных
	db := client.Database(databaseName)

	return &MongoClient{
		client: client,
		db:     db,
	}, nil
}

// GetCollection возвращает коллекцию по имени
func (mc *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return mc.db.Collection(collectionName)
}

// Close закрывает соединение с MongoDB
func (mc *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mc.client.Disconnect(ctx)
}
