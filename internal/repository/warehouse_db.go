package repository

import (
	"context"
	"errors"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type WarehouseRepositoryMongo struct {
	collection *mongo.Collection
}

// Конструктор репозитория для склада
func NewMongoWarehouseRepository(client *mongo.Client, dbName, collectionName string) *WarehouseRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &WarehouseRepositoryMongo{collection: collection}
}

// Получение склада по ID
func (repo *WarehouseRepositoryMongo) GetByID(id int) (*domain.WareHouse, error) {
	var warehouse domain.WareHouse
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&warehouse)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &warehouse, nil
}

// Создание нового склада
func (repo *WarehouseRepositoryMongo) Create(warehouse *domain.WareHouse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, warehouse)
	return err
}

// Обновление информации о складе
func (repo *WarehouseRepositoryMongo) Update(warehouse *domain.WareHouse) error {
	filter := bson.M{"id": warehouse.ID}
	update := bson.M{
		"$set": bson.M{
			"location": warehouse.Location,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("warehouse not found")
	}
	return nil
}

// Удаление склада
func (repo *WarehouseRepositoryMongo) Delete(id int) error {
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("warehouse not found")
	}
	return nil
}
