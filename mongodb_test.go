package authenticationMongo

import (
	"context"
	"github.com/mailhedgehog/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type fakePassword struct {
	pass string
	hash string
}

var fakePasswords = []fakePassword{
	{
		pass: "test1",
		hash: "$2a$12$CV3q6WzQBGEPqrPkh.hYn.HFO6mAxKfLLNxAMWIKx9wF93X6539nS",
	},
	{
		pass: "test2",
		hash: "$2a$12$6aBv1ox1kgMBcS9st4ixdu6HKW77DNdpyJNENN5vVMFqHHcF.q5Ra",
	},
}

func createMongoDbConnection() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetTimeout(5 * time.Second)

	clientOptions = clientOptions.SetAuth(options.Credential{
		Username: "test_root",
		Password: "test_secret",
	})

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	logger.PanicIfError(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	logger.PanicIfError(err)

	logManager().Debug("Connected to MongoDB")

	return client.Database("test_db")
}

func createMongoTestCollection() *mongo.Collection {
	collection := createMongoDbConnection().Collection("foo")

	// Truncate in case data saved form previous test.
	collection.DeleteMany(context.TODO(), bson.D{})

	_, err := collection.InsertOne(context.TODO(), UserRow{
		"default_user1",
		fakePasswords[0].hash,
		fakePasswords[1].hash,
		[]string{
			"1.1.1.1",
			"1.1.1.2",
		},
		[]string{
			"2.1.1.1",
			"2.1.1.2",
		},
		[]string{
			"foo@test.com",
			"bar@test.com",
		},
	})

	logger.PanicIfError(err)

	return collection
}
