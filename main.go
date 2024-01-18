package main

import (
	"context"
	"log"
	"os"

	"github.com/20-VIGNESH-K/crud_operations/config"
	"github.com/20-VIGNESH-K/crud_operations/routes"
	"github.com/20-VIGNESH-K/crud_operations/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initApp(mongoClient *mongo.Client) {
	routes.ProfileRoute(server)

}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	services.MongoClient = mongoclient
	services.Ctx = context.TODO()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	initApp(mongoclient)
	// fmt.Println("server running on port", config.Port)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(server.Run("0.0.0.0:" + port))

}
