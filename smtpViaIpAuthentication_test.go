package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestViaIpAuthentication(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaIpAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("fake", "1.1.1.1"))
}

func TestViaIpAuthenticationReturnsTrueIfNotEnabled(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaIpAuthentication.Enabled = false
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("fake", "1.1.1.1"))
}

func TestViaIpAuthenticationAddIp(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaIpAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().ViaIpAuthentication().AddIp("default_user1", "1.1.1.3"))

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))
}

func TestViaIpAuthenticationDeleteIp(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaIpAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().ViaIpAuthentication().DeleteIp("default_user1", "1.1.1.3"))
	(*gounit.T)(t).AssertNotError(auth.SMTP().ViaIpAuthentication().DeleteIp("default_user1", "1.1.1.1"))

	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))
}

func TestViaIpAuthenticationClearAllIps(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.ViaIpAuthentication.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().ViaIpAuthentication().ClearAllIps("default_user1"))

	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.1"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().ViaIpAuthentication().Authenticate("default_user1", "1.1.1.3"))
}
