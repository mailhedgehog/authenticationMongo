package authenticationMongo

import (
	"github.com/mailhedgehog/contracts"
	"github.com/mailhedgehog/gounit"
	"testing"
)

func TestIpsAllowListReturnsAlloweIfNotEnabled(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.IpsAllowList.Enabled = false
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("fake", "2.1.1.1"))
}

func TestIpsAllowList(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.IpsAllowList.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Enabled())

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("fake", "2.1.1.1"))
}

func TestIpsAllowListAddIp(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.IpsAllowList.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().IpsAllowList().AddIp("default_user1", "2.1.1.3"))

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))
}

func TestIpsAllowListDeleteIp(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.IpsAllowList.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().IpsAllowList().DeleteIp("default_user1", "2.1.1.5"))
	(*gounit.T)(t).AssertNotError(auth.SMTP().IpsAllowList().DeleteIp("default_user1", "2.1.1.1"))

	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.5"))
}

func TestIpsAllowListClearAllIps(t *testing.T) {
	config := &contracts.AuthenticationConfig{}
	config.Smtp.IpsAllowList.Enabled = true
	auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertTrue(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))

	(*gounit.T)(t).AssertNotError(auth.SMTP().IpsAllowList().ClearAllIps("default_user1"))

	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.1"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.2"))
	(*gounit.T)(t).AssertFalse(auth.SMTP().IpsAllowList().Allowed("default_user1", "2.1.1.3"))
}
