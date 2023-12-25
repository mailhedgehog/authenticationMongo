# MailHedgehog package to authenticate via MongoDB storage

All users data stored in MongoDB database. Useful if you have a lot of users in application.

## Usage

```go
package main

import (
    "github.com/mailhedgehog/gounit"
    "testing"
)

func Test(t *testing.T) {
    config := &contracts.AuthenticationConfig{}
    config.Smtp.ViaPasswordAuthentication.Enabled = true
    auth := CreateMongoDbAuthentication(createMongoTestCollection(), config)

    (*gounit.T)(t).AssertTrue(auth.SMTP().ViaPasswordAuthentication().Authenticate("user1", "secret"))
}
```

## Development

```shell
go mod tidy
go mod verify
go mod vendor
```

Test

```shell
docker-compose up -d
go test --cover
```

## Credits

- [![Think Studio](https://yaroslawww.github.io/images/sponsors/packages/logo-think-studio.png)](https://think.studio/)
