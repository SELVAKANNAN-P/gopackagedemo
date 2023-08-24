package controllers

import (
	"net/http"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"github.com/gin-gonic/gin"
)

type LoanTableController struct {
	LoanTableService interfaces.ILoanCollections 
}

func InitLoanTableController(LoanTableService interfaces.ILoanCollections) LoanTableController {
	return LoanTableController{LoanTableService: LoanTableService}
}

func (lc *LoanTableController) CreateLoanTable(ctx *gin.Context) {
	var loan models.LoanCollections
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result, err := lc.LoanTableService.CreateLoanTable(&loan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
}
