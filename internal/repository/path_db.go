package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PathRepositoryMongo struct {
	collection *mongo.Collection
}

// NewMongoPathRepository создает новый экземпляр репозитория для пути
func NewMongoPathRepository(collection *mongo.Collection) *PathRepositoryMongo {
	return &PathRepositoryMongo{collection: collection}
}

func (repo *PathRepositoryMongo) GetByID(id int) (*domain.Path, error) {
	var path domain.Path
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&path)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &path, err
}

func (repo *PathRepositoryMongo) Create(path *domain.Path) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, path)
	return err
}

func (repo *PathRepositoryMongo) Update(path *domain.Path) error {
	filter := bson.M{"id": path.StartLocationID}
	update := bson.M{
		"$set": bson.M{
			"endlocationid": path.EndLocationID,
			"distance":      path.Distance,
			"duration":      path.Duration,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *PathRepositoryMongo) Delete(path *domain.Path) error {
	filter := bson.M{"id": path.StartLocationID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
