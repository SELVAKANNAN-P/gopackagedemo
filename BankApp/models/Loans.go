package models

import "time"


type Loan struct{
	Id int64 `json:"customer_id" bson:"customer_id"`
	Name string `json:"name" bson:"name"`
	Amount int64 `json:"amount" bson:"amount"`
	Type string `json:"type" bson:"type"`
}
type sort struct{
	fromdate time.Time
	tdate time.Time
}