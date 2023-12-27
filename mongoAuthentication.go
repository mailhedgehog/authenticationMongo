package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	context *storageContext
}

type storageContext struct {
	collection *mongo.Collection
	config     *contracts.AuthenticationConfig
	storage    *Mongo
}

type UserRow struct {
	Username            string   `bson:"username"`
	DashboardPass       string   `bson:"dashboard_password"`
	SmtpPass            string   `bson:"smtp_password"`
	SmtpAuthIPs         []string `bson:"smtp_auth_ips"`
	SmtpAllowListedIPs  []string `bson:"smtp_allow_listed_ips"`
	DashboardAuthEmails []string `bson:"dashboard_auth_emails"`
}

func (mongo *Mongo) SMTP() contracts.SmtpAuthentication {
	return &smtpAuthentication{mongo.context}
}
func (mongo *Mongo) Dashboard() contracts.DashboardAuthentication {
	return &dashboardAuthentication{mongo.context}
}
func (mongo *Mongo) UsersStorage() contracts.UsersStorage {
	return &usersStorage{mongo.context}
}
