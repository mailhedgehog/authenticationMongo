package authenticationMongo

import (
	"context"
	"fmt"
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var configuredLogger *logger.Logger

func logManager() *logger.Logger {
	if configuredLogger == nil {
		configuredLogger = logger.CreateLogger("authenticationMongo")
	}
	return configuredLogger
}

func CreateMongoDbAuthentication(collection *mongo.Collection, config *contracts.AuthenticationConfig) *Mongo {
	indexName, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.M{"username": 1},
	})
	logger.PanicIfError(err)

	logManager().Debug(fmt.Sprintf("Index [%s] created", indexName))

	storage := &Mongo{
		context: &storageContext{
			collection: collection,
			config:     config,
		}}

	storage.context.storage = storage

	return storage
}
