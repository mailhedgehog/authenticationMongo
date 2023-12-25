package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestDashboardAuthenticateViaPassword(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaPasswordAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaPasswordAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[0].pass))
	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[1].pass))
	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", "fake"))

	// If user not exists returns false
	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaPasswordAuthentication().Authenticate("fake_not_exists", fakePasswords[0].pass))
}

func TestDashboardAuthenticateViaPasswordReturnsTrueIdNotEnabled(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaPasswordAuthentication.Enabled = false
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaPasswordAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", "fake"))
}

func TestDashboardAuthenticateViaPasswordStePass(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaPasswordAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[0].pass))

	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaPasswordAuthentication().SetPassword("default_user1", "new_foo_pass"))

	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", fakePasswords[0].pass))
	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaPasswordAuthentication().Authenticate("default_user1", "new_foo_pass"))

	(*gounit.T)(t).ExpectError(auth.Dashboard().ViaPasswordAuthentication().SetPassword("fake_not_exists", "foo"))
}
