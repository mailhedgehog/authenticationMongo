module github.com/mailhedgehog/authenticationMongo

go 1.19

require (
	github.com/mailhedgehog/contracts v1.0.0
	github.com/mailhedgehog/gounit v1.0.0
	github.com/mailhedgehog/logger v1.0.0
	go.mongodb.org/mongo-driver v1.13.1
	golang.org/x/crypto v0.17.0
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848
)

replace github.com/mailhedgehog/contracts v1.0.0 => ../contracts

require (
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/text v0.14.0 // indirect
)