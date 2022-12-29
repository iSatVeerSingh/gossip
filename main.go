package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iSatVeerSingh/gossip/db"
	"github.com/iSatVeerSingh/gossip/middlewares"
	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/routes"
	"github.com/iSatVeerSingh/gossip/utils"
)

// Load Environment Variables befor starting the application
func init() {
	utils.LoadEnv()
}

func main() {
	PORT := utils.GetEnv("PORT")

	mongoClient := db.GetMongoSession()
	defer db.MongoSessionClose(mongoClient)
	models.AddIndexes()

	// Intialize all routes
	router := routes.InitRoutes()

	router.Use(middlewares.Cors)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
