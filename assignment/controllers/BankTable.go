package controllers

import (
	"net/http"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"github.com/gin-gonic/gin"
)

type BankTableController struct {
	BankTableService interfaces.IBankTable
}

func InitBankTableController(BankTableService interfaces.IBankTable) BankTableController {
	return BankTableController{BankTableService: BankTableService}
}

func (bc *BankTableController) CreateBankTable(ctx *gin.Context) {
	var bankData models.BankTable
	if err := ctx.ShouldBindJSON(&bankData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	status, err := bc.BankTableService.CreateBankTable(&bankData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": status})
}
