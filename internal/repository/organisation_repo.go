package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
)

func (repo *MongoRepository) CreateOrganisation(organisation *model.Organisation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	organisation.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(ctx, organisation)
	return err
}

func (repo *MongoRepository) GetAllOrganisations() ([]model.Organisation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var organisations []model.Organisation
	filter := bson.M{"is_deleted": false} // Only fetch organisations where is_deleted is false

	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var organisation model.Organisation
		if err := cursor.Decode(&organisation); err != nil {
			return nil, err
		}
		organisations = append(organisations, organisation)
	}

	return organisations, nil
}
