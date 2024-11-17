package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LocationRepositoryMongo struct {
	collection *mongo.Collection
}

// Конструктор
func NewMongoLocationRepository(collection *mongo.Collection) *LocationRepositoryMongo {
	return &LocationRepositoryMongo{collection: collection}
}

func (repo *LocationRepositoryMongo) GetByID(id int) (domain.Location, error) {
	var location domain.Location
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&location)
	if err == mongo.ErrNoDocuments {
		return domain.Location{}, nil // Возвращаем пустую структуру, если документ не найден
	}
	return location, err
}

func (repo *LocationRepositoryMongo) Create(location *domain.Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, location)
	return err
}

func (repo *LocationRepositoryMongo) Update(location *domain.Location) error {
	filter := bson.M{"id": location.ID}
	update := bson.M{
		"$set": bson.M{
			"name":      location.Name,
			"address":   location.Address,
			"latitude":  location.Latitude,
			"longitude": location.Longitude,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *LocationRepositoryMongo) Delete(id int) error {
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
