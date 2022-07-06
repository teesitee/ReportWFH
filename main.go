package main

import (
	"context"

	"report-lkl-morning/api"
	"report-lkl-morning/repository"
	"report-lkl-morning/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ReportInfo := client.Database("ReportWFH").Collection("InfoWFH")
	rp := repository.MongoRepository{
		Collection: ReportInfo,
	}
	h := api.Handler{
		Repo: rp,
	}

	r := router.NewRouter(h)
	errRun := r.Run(":8080")
	if errRun != nil {
		panic(errRun)
	}

}
