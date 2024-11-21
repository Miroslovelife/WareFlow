package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type GenericRepository[T any] interface {
	GetByID(id int) (*T, error)
	Create(entity *T) error
	Update(id int, updates bson.M) error
	Delete(id int) error
}

type MongoGenericRepository[T any] struct {
	collection *mongo.Collection
}

func NewMongoGenericRepository[T any](collection *mongo.Collection) *MongoGenericRepository[T] {
	return &MongoGenericRepository[T]{collection: collection}
}

func (repo *MongoGenericRepository[T]) GetByID(id int) (*T, error) {
	var entity T
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&entity)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *MongoGenericRepository[T]) Create(entity *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, entity)
	return err
}

func (repo *MongoGenericRepository[T]) Update(id int, updates bson.M) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": updates}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("entity not found")
	}
	return nil
}

func (repo *MongoGenericRepository[T]) Delete(id int) error {
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("entity not found")
	}
	return nil
}
