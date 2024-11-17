package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type WarehouseRepositoryMongo struct {
	collection *mongo.Collection
}

// Конструктор
func NewMongoWarehouseRepository(collection *mongo.Collection) *WarehouseRepositoryMongo {
	return &WarehouseRepositoryMongo{collection: collection}
}

func (repo *WarehouseRepositoryMongo) GetByID(id int) (*domain.WareHouse, error) {
	var warehouse domain.WareHouse
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&warehouse)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &warehouse, err
}

func (repo *WarehouseRepositoryMongo) Create(warehouse *domain.WareHouse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, warehouse)
	return err
}

func (repo *WarehouseRepositoryMongo) Update(warehouse *domain.WareHouse) error {
	filter := bson.M{"id": warehouse.ID}
	update := bson.M{
		"$set": bson.M{
			"location": warehouse.Location,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *WarehouseRepositoryMongo) Delete(id int) error {
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
