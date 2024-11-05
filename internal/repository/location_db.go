package repository

import (
	"WareFlow/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepositoryMongo struct {
	collection *mongo.Collection
}

func NewMongoLocationRepository(client *mongo.Client, dbName, collectionName string) *LocationRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &LocationRepositoryMongo{collection: collection}
}

func (repo *LocationRepositoryMongo) GetByID(id int) (*domain.Location, error) {
	var location domain.Location
	filter := bson.M{"id": id}

	err := repo.collection.FindOne(context.TODO(), filter).Decode(&location)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &location, nil
}

func (repo *LocationRepositoryMongo) Create(location *domain.Location) error {
	_, err := repo.collection.InsertOne(context.TODO(), location)
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

	_, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (repo *LocationRepositoryMongo) Delete(location *domain.Location) error {
	filter := bson.M{"id": location.ID}

	_, err := repo.collection.DeleteOne(context.TODO(), filter)
	return err
}
