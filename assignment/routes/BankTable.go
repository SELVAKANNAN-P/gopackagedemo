package routes

import (
	"rest-api-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func BankTableRoute(router *gin.Engine, controllers controllers.BankTableController) {
	router.POST("/api/BankTable/create", controllers. CreateBankTable)
	// router.GET("/api/Transaction/:id", controllers.GetTransactionById)
}
