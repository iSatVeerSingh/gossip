package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iSatVeerSingh/gossip/db"
	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/routes"
	"github.com/iSatVeerSingh/gossip/utils"
)

func init() {
	utils.LoadEnv()
}

func main() {
	PORT := utils.GetEnv("PORT")

	mongoClient := db.GetMongoSession()
	defer db.MongoSessionClose(mongoClient)
	models.AddIndexes()

	router := routes.InitRoutes()

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
