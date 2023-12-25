package authenticationMongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
)

type smtpViaIpAuthentication struct {
}

func (authentication *smtpViaIpAuthentication) Enabled() bool {
	return mongoClient.config.Smtp.ViaIpAuthentication.Enabled
}

func (authentication *smtpViaIpAuthentication) Authenticate(username string, ip string) bool {
	if !authentication.Enabled() {
		return true
	}

	var user UserRow
	err := mongoClient.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	return slices.Contains(user.SmtpAuthIPs, ip)
}

func (authentication *smtpViaIpAuthentication) AddIp(username string, ip string) error {
	var user UserRow
	err := mongoClient.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	if slices.Contains(user.SmtpAuthIPs, ip) {
		return nil
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	user.SmtpAuthIPs = append(user.SmtpAuthIPs, ip)

	newValues = append(newValues, bson.E{"smtp_auth_ips", user.SmtpAuthIPs})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	if updateResult.MatchedCount <= 0 {
		err = errors.New("user with such username not found")
	}

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (authentication *smtpViaIpAuthentication) DeleteIp(username string, ip string) error {
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

	if !slices.Contains(user.SmtpAuthIPs, ip) {
		return nil
	}

	i := slices.Index(user.SmtpAuthIPs, ip)
	user.SmtpAuthIPs = slices.Delete(user.SmtpAuthIPs, i, i+1)

	newValues = append(newValues, bson.E{"smtp_auth_ips", user.SmtpAuthIPs})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (authentication *smtpViaIpAuthentication) ClearAllIps(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	newValues = append(newValues, bson.E{"smtp_auth_ips", []string{}})

	updateResult, err := mongoClient.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}
