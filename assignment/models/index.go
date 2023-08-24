package models

type CustomerTable struct {
	ID       int    `json:"cust_id" bson:"cust_id"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

type CustomerTransactions struct {
	CustomerID        string  `bson:"customer_id"`
	TransactionAmount float64 `bson:"transaction_amount"`
}

type LoanCollections struct {
	CustomerID string  `bson:"cus_id"`
	LoanAmount float64 `bson:"loan_amount"`
	LoanType   string  `bson:"loan_type"`
}

type AccountsTable struct {
	ID          int    `bson:"_acc_id"`
	CustomerID  int    `bson:"cus_id"`
	AccountType string `bson:"acc_type"`
	BranchAddr  string `bson:"acc_branch_address"`
}

type BankTable struct {
	ID          string `json:"bank_id" bson:"bank_id"`
	IFSCCode    string `json:"ifsc_code" bson:"ifsc_code"`
	BankAddress string `bson:"bank_address"`
}
type UpdatDetails struct {
	OldName string `bson:"oldname"`
	NewName string `bson:"newname"`
}
type DeleteCustomerTable struct {
	IFSCCode int64 `json:"ifsc_code" bson:"ifsc_code"`
}
