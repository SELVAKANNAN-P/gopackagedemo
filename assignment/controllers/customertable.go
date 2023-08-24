package controllers

import (
	"net/http"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"github.com/gin-gonic/gin"
)

type CustomerTableController struct {
	CustomerTableService interfaces.ICustomerTable
}

func InitCustomerTableontroller(CustomerTableService interfaces.ICustomerTable) CustomerTableController {

	return CustomerTableController{CustomerTableService}
}

func (pc *CustomerTableController) CreateCustomerTable(ctx *gin.Context) {

	var profile *models.CustomerTable
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.CustomerTableService.CreateCustomerTable(profile)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}

func (pc *CustomerTableController) UpdateAccountsTable(ctx *gin.Context) {
	var profile *models.UpdatDetails
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	customer, err := pc.CustomerTableService.UpdateCustomerTable(profile.OldName, profile.NewName)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": customer})

}
func (cc *CustomerTableController) DeleteAccountsTable(ctx *gin.Context) {
	// name := ctx.Param("cust_id")
	result, err := cc.CustomerTableService.DeleteCustomerTable(9)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}
