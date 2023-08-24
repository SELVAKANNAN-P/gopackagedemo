package controllers

import (
	"bankDemo/interfaces"
	"bankDemo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountService interfaces.Iaccount
}

func InitAccController(accountService interfaces.Iaccount) AccountController {
	return AccountController{accountService}
}

func (a *AccountController) CreateAccount(ctx *gin.Context) {
	var acc *models.Account
	if err := ctx.ShouldBindJSON(&acc); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newacc, err := a.AccountService.CreateAccount(acc)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newacc})
}

func (a *AccountController) GetAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, _ := strconv.ParseInt(id, 10, 64)
	acc, err := a.AccountService.GetAccountById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": acc})
}
func (a *AccountController) UpdateAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	acc := &models.Account{}
	if err := ctx.ShouldBindJSON(&acc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid ID format"})
		return
	}

	// Update only if there's something to update (not all fields are required)
	// if acc.Account_type != "" || acc.Branch != "" {
	index, err := a.AccountService.UpdateAccountById(id1, acc)
	index = index
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Account updated successfully"})
}

func (a *AccountController) DeleteAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, _ := strconv.ParseInt(id, 10, 64)
	acc, err := a.AccountService.DeleteAccountById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": acc})
}
