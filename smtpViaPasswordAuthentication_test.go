package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestSmtpAuthenticateViaPassword(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaPasswordAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[1].pass))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[0].pass))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", "fake"))

	// If user not exists returns false
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaPasswordAuthentication().Authenticate("fake_not_exists", fakePasswords[0].pass))
}

func TestSmtpAuthenticateViaPasswordReturnsTrueIdNotEnabled(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaPasswordAuthentication.Enabled = false
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaPasswordAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", "fake"))
}

func TestSmtpAuthenticateViaPasswordStePass(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaPasswordAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[1].pass))

	(*gounit.T)(t).AssertNotError(auth.SMTP().ViaPasswordAuthentication().SetPassword("default_user1", "new_foo_pass"))

	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[1].pass))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Authenticate("default_user1", "new_foo_pass"))

	(*gounit.T)(t).ExpectError(auth.SMTP().ViaPasswordAuthentication().SetPassword("fake_not_exists", "foo"))
}
