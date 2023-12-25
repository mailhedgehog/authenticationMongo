package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestUsersStorageExists(t *testing.T) {
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), &contracts.AuthenticationConfig{})

	(*gounit.T)(t).AssertTrue(auth.UsersStorage().Exists("default_user1"))
	(*gounit.T)(t).AssertFalse(auth.UsersStorage().Exists("user0"))

	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user0"))

	(*gounit.T)(t).AssertTrue(auth.UsersStorage().Exists("default_user1"))
	(*gounit.T)(t).AssertTrue(auth.UsersStorage().Exists("user0"))
}

func TestUsersStorageList(t *testing.T) {
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), &contracts.AuthenticationConfig{})

	list, foundCount, err := auth.UsersStorage().List("user", 0, 100)

	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(1, foundCount)
	(*gounit.T)(t).AssertEqualsInt(1, len(list))

	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user1"))
	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user2"))
	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user3"))

	list, foundCount, err = auth.UsersStorage().List("user", 1, 100)

	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(4, foundCount)
	(*gounit.T)(t).AssertEqualsInt(3, len(list))
}

func TestUsersStorageDelete(t *testing.T) {
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), &contracts.AuthenticationConfig{})

	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user1"))
	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user2"))
	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user3"))
	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Add("user4"))

	_, foundCount, err := auth.UsersStorage().List("user", 0, 100)

	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(5, foundCount)

	(*gounit.T)(t).AssertTrue(auth.UsersStorage().Exists("user3"))

	(*gounit.T)(t).AssertNotError(auth.UsersStorage().Delete("user3"))

	(*gounit.T)(t).AssertFalse(auth.UsersStorage().Exists("user3"))

	_, foundCount, err = auth.UsersStorage().List("user", 0, 100)

	(*gounit.T)(t).AssertNotError(err)
	(*gounit.T)(t).AssertEqualsInt(4, foundCount)
}
