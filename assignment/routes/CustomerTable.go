package routes

import (
	"rest-api-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerTableRoute(router *gin.Engine, controllers controllers.CustomerTableController) {
	// router.POST("/api/customerTable/create", controllers.CreateCustomerTable)
 	// // router.GET("/api/Transaction/:id", controllers.GetTransactionById)
	router.POST("/api/updatecustomer",controllers.UpdateAccountsTable)
    //router.DELETE("/api/customerTable/delete/:cust_id", controllers.DeleteAccountsTable)
 }
//  func BankTableRoute(router *gin.Engine, controllers controllers.BankTableController){
// 	router.POST("/api/BankTable/create", controllers.CreateTCustomerTable)

//  }
