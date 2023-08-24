package services

import (
	"context"
	"fmt"
	"rest-api-mongodb/interfaces"
	"rest-api-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerTableService struct {
	CustomerTableCollection *mongo.Collection
	ctx                     context.Context
}

func NewCustomerServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomerTable {
	return &CustomerTableService{
		CustomerTableCollection: collection,
		ctx:                     ctx,
	}
}

func (c *CustomerTableService) CreateCustomerTable(user *models.CustomerTable) (string, error) {
	_, err := c.CustomerTableCollection.InsertOne(c.ctx, user)
	if err != nil {
		return "error", err
	}
	return "success", nil
}

func (c *CustomerTableService) UpdateCustomerTable(oldname string, newname string) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "name", Value: oldname}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newname}}}}
	result, err := c.CustomerTableCollection.UpdateOne(c.ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func (c *CustomerTableService) DeleteCustomerTable(cust_id int) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "cust_id", Value: cust_id}}
	result, err := c.CustomerTableCollection.DeleteOne(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
