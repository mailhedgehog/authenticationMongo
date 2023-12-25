package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestViaEmailAuthenticationSendToken(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaEmailAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	// TODO: not implemented
	(*gounit.T)(t).ExpectError(auth.Dashboard().ViaEmailAuthentication().SendToken("user1", "test@test.com"))
}

func TestViaEmailAuthenticationAuthenticate(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaEmailAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	// TODO: not implemented
	(*gounit.T)(t).AssertFalse(auth.Dashboard().ViaEmailAuthentication().Authenticate("user1", "test@test.com", "123"))
}

func TestViaEmailAuthenticationAddEmail(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaEmailAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaEmailAuthentication().Enabled())

	users, foundCount, err := auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(2, len(users[0].DashboardAuthEmails))

	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaEmailAuthentication().AddEmail("default_user1", "foo@test.com"))
	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaEmailAuthentication().AddEmail("default_user1", "test@test.com"))

	users, foundCount, err = auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(3, len(users[0].DashboardAuthEmails))
}

func TestViaEmailAuthenticationDeleteEmail(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaEmailAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaEmailAuthentication().Enabled())

	users, foundCount, err := auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(2, len(users[0].DashboardAuthEmails))

	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaEmailAuthentication().DeleteEmail("default_user1", "foo@test.com"))
	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaEmailAuthentication().DeleteEmail("default_user1", "test@test.com"))

	users, foundCount, err = auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(1, len(users[0].DashboardAuthEmails))
	(*gounit.T)(t).AssertEqualsString("bar@test.com", users[0].DashboardAuthEmails[0])
}

func TestViaEmailAuthenticationClearAllEmails(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Dashboard.ViaEmailAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.Dashboard().ViaEmailAuthentication().Enabled())

	users, foundCount, err := auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(2, len(users[0].DashboardAuthEmails))

	(*gounit.T)(t).AssertNotError(auth.Dashboard().ViaEmailAuthentication().ClearAllEmails("default_user1"))

	users, foundCount, err = auth.UsersStorage().List("user", 0, 100)
	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(0, len(users[0].DashboardAuthEmails))
}
