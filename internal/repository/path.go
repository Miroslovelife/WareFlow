package repository

import (
	"WareFlow/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PathRepositoryMongo struct {
	collection *mongo.Collection
}

func NewMongoPathRepository(client *mongo.Client, dbName, collectionName string) *PathRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &PathRepositoryMongo{collection: collection}
}

func (repo *PathRepositoryMongo) GetByID(id int) (*domain.Path, error) {
	var path domain.Path
	filter := bson.M{"id": id}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&path)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &path, nil
}

func (repo *PathRepositoryMongo) Create(path *domain.Path) error {
	_, err := repo.collection.InsertOne(context.TODO(), path)
	return err
}

func (repo *PathRepositoryMongo) Update(path *domain.Path) error {
	filter := bson.M{"id": path.StartLocationID} // Используем `StartLocationID` или другой уникальный идентификатор записи
	update := bson.M{
		"$set": bson.M{
			"endlocationid": path.EndLocationID,
			"distance":      path.Distance,
			"duration":      path.Duration,
		},
	}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (repo *PathRepositoryMongo) Delete(path *domain.Path) error {
	filter := bson.M{"id": path.StartLocationID} // Используем `StartLocationID` или другой уникальный идентификатор записи

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	return err
}
