package services

import (
	"context"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoanTableService struct {
	LoanTableCollection *mongo.Collection
	ctx                 context.Context
}

// CreateBankTable implements interfaces.IBankTable.
func (*LoanTableService) CreateBankTable(Bank *models.BankTable) (string, error) {
	panic("unimplemented")
}

func NewLoanServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ILoanCollections  {
	return &LoanTableService{
		LoanTableCollection: collection,
		ctx:                 ctx,
	}
}

func (c *LoanTableService) CreateLoanTable(loan *models.LoanCollections) (string, error) {
	_, err := c.LoanTableCollection.InsertOne(c.ctx, loan)
	if err != nil {
		return "error", err
	}
	return "success", nil
}
