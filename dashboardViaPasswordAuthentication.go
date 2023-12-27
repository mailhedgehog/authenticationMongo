package authenticationMongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailhedgehog/contracts"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type dashboardViaPasswordAuthentication struct {
	context *storageContext
}

func (authentication *dashboardViaPasswordAuthentication) Enabled() bool {
	return authentication.context.config.Dashboard.ViaPasswordAuthentication.Enabled
}

func (authentication *dashboardViaPasswordAuthentication) Authenticate(username string, password string) bool {
	if !authentication.Enabled() {
		return true
	}

	var user UserRow
	err := authentication.context.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.DashboardPass), []byte(password)); err != nil {
		return false
	}

	return true
}

func (authentication *dashboardViaPasswordAuthentication) SetPassword(username string, password string) error {
	if len(username) <= 0 {
		return errors.New("username required")
	}

	var newPassHash []byte
	if len(password) > 0 {
		var err error
		newPassHash, err = contracts.CreatePasswordHash(password)
		if err != nil {
			return err
		}
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	newValues = append(newValues, bson.E{"dashboard_password", newPassHash})

	updateResult, err := authentication.context.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	if updateResult.MatchedCount <= 0 {
		err = errors.New("user with such username not found")
	}

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}
