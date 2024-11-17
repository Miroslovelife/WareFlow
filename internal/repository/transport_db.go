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

func NewTransportRepositoryMongo(db *mongo.Database) *TransportRepositoryMongo {
	return &TransportRepositoryMongo{
		collection: db.Collection("transports"),
	}
}

func (r *TransportRepositoryMongo) GetByID(id int) (*domain.Transport, error) {
	var transport domain.Transport
	filter := bson.M{"id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&transport)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &transport, nil
}

func (r *TransportRepositoryMongo) Create(transport *domain.Transport) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, transport)
	return err
}

func (r *TransportRepositoryMongo) Update(transport *domain.Transport) error {
	filter := bson.M{"id": transport.ID}
	update := bson.M{
		"$set": bson.M{
			"type":           transport.Type,
			"capacity":       transport.CapacityVolume,
			"capacityWeight": transport.CapacityWeight,
			"expense":        transport.Expense,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("transport not found")
	}
	return nil
}

func (r *TransportRepositoryMongo) Delete(id int) error {
	filter := bson.M{"id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("transport not found")
	}
	return nil
}
