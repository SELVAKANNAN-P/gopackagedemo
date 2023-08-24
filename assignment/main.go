package main

import (
	"context"
	"fmt"
	"log"
	"rest-api-mongodb/config"
	"rest-api-mongodb/constants"
	"rest-api-mongodb/controllers"
	"rest-api-mongodb/routes"
	"rest-api-mongodb/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}
func initcApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	//create collection
	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("transaction")
	//pass the collection and context to initantiate the service
	profileService := services.NewCustomerServiceInit(profileCollection, ctx)
	//pass the service to the controller
	profileController := controllers.InitCustomerTableontroller(profileService)
	//pass the controller to the routes
	routes.CustomerTableRoute(server, profileController)

}
func initbApp(mongoClient *mongo.Client) {
	ctx := context.TODO()
	bankTableCollection := mongoClient.Database(constants.DatabaseName).Collection("bankTable")
	bankTableService := services.NewTableServiceInit(bankTableCollection, ctx)
	bankTableController := controllers.InitBankTableController(bankTableService)

	routes.BankTableRoute(server, bankTableController)
}
func iniltApp(mongoClient *mongo.Client) {
	ctx := context.TODO()
	LoanCollections := mongoClient.Database(constants.DatabaseName).Collection("LoanTable ")
	LoanServices := services.NewLoanServiceInit(LoanCollections, ctx)
	LoanControllers := controllers.InitLoanTableController(LoanServices)
	routes.LoanTableRoute(server, LoanControllers)
}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initcApp(mongoclient)

	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
