package repository

import (
	"context"
	"errors"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransportRepositoryMongo struct {
	collection *mongo.Collection
}

// NewTransportRepositoryMongo создает репозиторий для транспорта
func NewTransportRepositoryMongo(collection *mongo.Collection) *TransportRepositoryMongo {
	return &TransportRepositoryMongo{collection: collection}
}

func (repo *TransportRepositoryMongo) GetByID(id int) (*domain.Transport, error) {
	var transport domain.Transport
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&transport)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &transport, err
}

func (repo *TransportRepositoryMongo) Create(transport *domain.Transport) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, transport)
	return err
}

func (repo *TransportRepositoryMongo) Update(transport *domain.Transport) error {
	filter := bson.M{"id": transport.ID}
	update := bson.M{
		"$set": bson.M{
			"type":           transport.Type,
			"capacityvolume": transport.CapacityVolume,
			"capacityweight": transport.CapacityWeight,
			"expense":        transport.Expense,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("transport not found")
	}
	return nil
}

func (repo *TransportRepositoryMongo) Delete(id int) error {
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("transport not found")
	}
	return nil
}
