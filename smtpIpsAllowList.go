package authenticationMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
)

type smtpIpsAllowList struct {
	context *storageContext
}

func (allowlist *smtpIpsAllowList) Enabled() bool {
	return allowlist.context.config.Smtp.IpsAllowList.Enabled
}

func (allowlist *smtpIpsAllowList) Allowed(username string, ip string) bool {
	if !allowlist.Enabled() {
		return true
	}

	var user UserRow
	err := allowlist.context.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		logManager().Debug(err.Error())
		return false
	}

	return slices.Contains(user.SmtpAllowListedIPs, ip)
}

func (allowlist *smtpIpsAllowList) AddIp(username string, ip string) error {
	var user UserRow
	err := allowlist.context.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return err
	}

	if slices.Contains(user.SmtpAllowListedIPs, ip) {
		return nil
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	user.SmtpAllowListedIPs = append(user.SmtpAllowListedIPs, ip)

	newValues = append(newValues, bson.E{"smtp_allow_listed_ips", user.SmtpAllowListedIPs})

	updateResult, err := allowlist.context.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (allowlist *smtpIpsAllowList) DeleteIp(username string, ip string) error {
	var user UserRow
	err := allowlist.context.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
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

	if !slices.Contains(user.SmtpAllowListedIPs, ip) {
		return nil
	}

	i := slices.Index(user.SmtpAllowListedIPs, ip)
	user.SmtpAllowListedIPs = slices.Delete(user.SmtpAllowListedIPs, i, i+1)

	newValues = append(newValues, bson.E{"smtp_allow_listed_ips", user.SmtpAllowListedIPs})

	updateResult, err := allowlist.context.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}

func (allowlist *smtpIpsAllowList) ClearAllIps(username string) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"username", username}},
			}},
	}

	newValues := bson.D{}

	newValues = append(newValues, bson.E{"smtp_allow_listed_ips", []string{}})

	updateResult, err := allowlist.context.collection.UpdateOne(context.TODO(), filter, bson.D{bson.E{"$set", newValues}})

	logManager().Debug(fmt.Sprintf("User [%s] updated, mongo _id='%s'", username, updateResult.UpsertedID))

	return err
}
