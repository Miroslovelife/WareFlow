package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OptimizationResultRepositoryMongo struct {
	collection *mongo.Collection
}

// NewMongoOptimizationResultRepository создает репозиторий для результатов оптимизации
func NewMongoOptimizationResultRepository(collection *mongo.Collection) *OptimizationResultRepositoryMongo {
	return &OptimizationResultRepositoryMongo{collection: collection}
}

func (repo *OptimizationResultRepositoryMongo) GetByID(id int) (*domain.OptimizationResult, error) {
	var result domain.OptimizationResult
	filter := bson.M{"id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}

func (repo *OptimizationResultRepositoryMongo) Create(result *domain.OptimizationResult) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, result)
	return err
}

func (repo *OptimizationResultRepositoryMongo) Update(result *domain.OptimizationResult) error {
	filter := bson.M{"id": result.TransportID}
	update := bson.M{
		"$set": bson.M{
			"totaldistance": result.TotalDistance,
			"totalcost":     result.TotalCost,
			"route":         result.Route,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *OptimizationResultRepositoryMongo) Delete(result *domain.OptimizationResult) error {
	filter := bson.M{"id": result.TransportID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, filter)
	return err
}
