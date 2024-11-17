package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CargoRepositoryMongo struct {
	collection *mongo.Collection
}

// Конструктор репозитория для Cargo
func NewMongoCargoRepository(collection *mongo.Collection) *CargoRepositoryMongo {
	return &CargoRepositoryMongo{collection: collection}
}

func (repo *CargoRepositoryMongo) GetByID(id int) (*domain.Cargo, error) {
	var cargo domain.Cargo
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&cargo)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &cargo, err
}

func (repo *CargoRepositoryMongo) Create(cargo *domain.Cargo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, cargo)
	return err
}

func (repo *CargoRepositoryMongo) Update(cargo *domain.Cargo) error {
	filter := bson.M{"id": cargo.ID}
	update := bson.M{
		"$set": bson.M{
			"weight":      cargo.Weight,
			"volume":      cargo.Volume,
			"description": cargo.Description,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *CargoRepositoryMongo) Delete(cargo *domain.Cargo) error {
	filter := bson.M{"id": cargo.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
