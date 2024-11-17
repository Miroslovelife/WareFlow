package repository

import (
	"context"
	"github.com/Miroslovelife/WareFlow/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OptimizationResultRepositoryMongo struct {
	collection *mongo.Collection
}

func NewMongoOptimizationResultRepository(client *mongo.Client, dbName, collectionName string) *OptimizationResultRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &OptimizationResultRepositoryMongo{collection: collection}
}

func (repo *OptimizationResultRepositoryMongo) GetByID(id int) (*domain.OptimizationResult, error) {
	var result domain.OptimizationResult
	filter := bson.M{"transportid": id}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (repo *OptimizationResultRepositoryMongo) Create(result *domain.OptimizationResult) error {
	_, err := repo.collection.InsertOne(context.TODO(), result)
	return err
}

func (repo *OptimizationResultRepositoryMongo) Update(result *domain.OptimizationResult) error {
	filter := bson.M{"transportid": result.TransportID}
	update := bson.M{
		"$set": bson.M{
			"totaldistance": result.TotalDistance,
			"totalcost":     result.TotalCost,
			"route":         result.Route,
		},
	}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (repo *OptimizationResultRepositoryMongo) Delete(result *domain.OptimizationResult) error {
	filter := bson.M{"transportid": result.TransportID}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	return err
}
