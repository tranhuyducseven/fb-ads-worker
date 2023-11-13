package main

import (
	"crypto/tls"
	"crypto/x509"
	job "data-pipeline/job"
	"fmt"
	"github.com/ponlv/go-kit/mongodb"
	"github.com/ponlv/go-kit/postgresql"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load(os.ExpandEnv("/config/.env"))
	if err != nil {
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			fmt.Println(err)
		}
	}

	// load postgresql certificate
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(os.Getenv("POSTGRESQL_CERT")))
	if !ok {
		panic("failed to parse root certificate")
	}

	// init postgresql
	postgresql.InitDB(postgresql.Config{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT")),
		User:     os.Getenv("POSTGRESQL_USERNAME"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
		TLSConfig: &tls.Config{
			RootCAs:    roots,
			ServerName: os.Getenv("POSTGRESQL_HOST"),
		},
	})

	// init mongodb
	mongoConfig := &mongodb.MongoConfig{
		Host:              os.Getenv("MONGODB_HOST"),
		Username:          os.Getenv("MONGODB_USERNAME"),
		Password:          os.Getenv("MONGODB_PASSWORD"),
		UseSRV:            true,
		DbName:            os.Getenv("MONGODB_DATABASE"),
		MaxConnectionPool: 1000,
	}
	_, _, _, err = mongodb.ConnectMongoWithConfig(mongoConfig, nil, nil)
	if err != nil {
		panic(err)
	}

	job.RunAllJob()

}
