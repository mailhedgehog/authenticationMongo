package authenticationMongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailhedgehog/contracts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type usersStorage struct {
	context *storageContext
}

func (storage *usersStorage) Exists(username string) bool {
	count, err := storage.context.collection.CountDocuments(context.TODO(), bson.M{"username": username})

	if err != nil || count <= 0 {
		return false
	}

	return true
}

func (storage *usersStorage) Add(username string) error {
	insertResult, err := storage.context.collection.InsertOne(context.TODO(), UserRow{
		username,
		"",
		"",
		[]string{},
		[]string{},
		[]string{},
	})

	logManager().Debug(fmt.Sprintf("New user [%s] added, mongo _id='%s'", username, insertResult.InsertedID))

	return err
}

func (storage *usersStorage) Delete(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	result, err := storage.context.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New(fmt.Sprintf("Unexpected count of deleted items, extected 1, got %d", result.DeletedCount))
	}

	return nil
}

func (storage *usersStorage) List(searchQuery string, offset, limit int) ([]contracts.UserResource, int, error) {
	opts := options.Find().SetSort(bson.M{"username": 1}).SetSkip(int64(offset)).SetLimit(int64(limit))
	textsMatch := bson.A{
		bson.M{"username": primitive.Regex{Pattern: searchQuery, Options: ""}},
	}
	filterQuery := bson.A{}
	if len(textsMatch) > 0 {
		filterQuery = append(filterQuery, bson.M{"$or": textsMatch})
	}
	filter := bson.M{"$and": filterQuery}

	totalCount, err := storage.context.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := storage.context.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, 0, err
	}
	var resources []contracts.UserResource
	var results []UserRow
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, 0, err
	}
	for _, result := range results {
		resources = append(resources, contracts.UserResource{
			Username:            result.Username,
			SmtpAuthIPs:         result.SmtpAuthIPs,
			SmtpAllowListedIPs:  result.SmtpAllowListedIPs,
			DashboardAuthEmails: result.DashboardAuthEmails,
		})
	}

	return resources, int(totalCount), nil
}
