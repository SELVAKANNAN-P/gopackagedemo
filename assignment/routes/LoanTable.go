package routes

import (
	"rest-api-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func LoanTableRoute(router *gin.Engine, controllers controllers.LoanTableController) {
	router.POST("/api/LoanTable/create", controllers.CreateLoanTable)
	// router.GET("/api/Transaction/:id", controllers.GetTransactionById)
}
