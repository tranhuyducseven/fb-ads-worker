package main

import (
	job "data-pipeline/job"
	"fmt"
	"github.com/ponlv/go-kit/mongodb"
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
