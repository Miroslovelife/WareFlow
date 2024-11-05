package repository

import (
	"WareFlow/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CargoRepositoryMongo struct {
	collection *mongo.Collection
}

func NewMongoCargoRepository(client *mongo.Client, dbName, collectionName string) *CargoRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &CargoRepositoryMongo{collection: collection}
}

func (repo *CargoRepositoryMongo) GetByID(id int) (*domain.Cargo, error) {
	var cargo domain.Cargo
	filter := bson.M{"id": id}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&cargo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &cargo, nil
}

func (repo *CargoRepositoryMongo) Create(cargo *domain.Cargo) error {
	_, err := repo.collection.InsertOne(context.TODO(), cargo)
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

	_, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (repo *CargoRepositoryMongo) Delete(cargo *domain.Cargo) error {
	filter := bson.M{"id": cargo.ID}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	return err
}
