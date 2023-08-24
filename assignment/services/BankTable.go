package services

import (
	"context"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type BankTableService struct {
	BankTableCollection *mongo.Collection
	ctx                 context.Context
}

func NewTableServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IBankTable {
	return &BankTableService{
		BankTableCollection: collection,
		ctx:                 ctx,
	}
}

func (t *BankTableService) CreateBankTable(user *models.BankTable) (string, error) {
	_, err := t.BankTableCollection.InsertOne(t.ctx, user)
	if err != nil {
		return "error", err
	}
	return "success", nil
}
