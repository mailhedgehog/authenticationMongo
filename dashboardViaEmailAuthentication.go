package authenticationMongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
)

type dashboardViaEmailAuthentication struct {
}

func (authentication *dashboardViaEmailAuthentication) Enabled() bool {
	return mongoClient.config.Dashboard.ViaEmailAuthentication.Enabled
}

func (authentication *dashboardViaEmailAuthentication) SendToken(username string, email string) error {
	return errors.New("functionality not implemented. please request developer to implement")
}

func (authentication *dashboardViaEmailAuthentication) Authenticate(username string, email string, token string) bool {
	return false
}

func (authentication *dashboardViaEmailAuthentication) AddEmail(username string, email string) error {
	var user UserRow
	err := mongoClient.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	if slices.Contains(user.DashboardAuthEmails, email) {
		return nil
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	user.DashboardAuthEmails = append(user.DashboardAuthEmails, email)

	newValues = append(newValues, bson.E{"dashboard_auth_emails", user.DashboardAuthEmails})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (authentication *dashboardViaEmailAuthentication) DeleteEmail(username string, email string) error {
	var user UserRow
	err := mongoClient.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	if !slices.Contains(user.DashboardAuthEmails, email) {
		return nil
	}

	i := slices.Index(user.DashboardAuthEmails, email)
	user.DashboardAuthEmails = slices.Delete(user.DashboardAuthEmails, i, i+1)

	newValues = append(newValues, bson.E{"dashboard_auth_emails", user.DashboardAuthEmails})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (authentication *dashboardViaEmailAuthentication) ClearAllEmails(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	newValues = append(newValues, bson.E{"dashboard_auth_emails", []string{}})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}
