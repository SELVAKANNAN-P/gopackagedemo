package interfaces

import (
	"rest-api-mongodb/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICustomerTable interface {
	CreateCustomerTable(Customer *models.CustomerTable) (string, error)
	UpdateCustomerTable(oldname string, newname string) (*mongo.UpdateResult, error)
	DeleteCustomerTable(cust_id int) (*mongo.DeleteResult, error)
	// GetCustomerTable(id string) (string, error)
}

// type ICustomerTransactions interface {
// 	CreateCustomerTransactions(CustomerTransactions *models.CustomerTransactions) (string, error)
// 	UpdateCustomerTransactions(id string) (string, error)
// 	DeleteCustomerTransactions(CustomerTransactions *models.CustomerTransactions) (string, error)
// 	GetCustomerTransactions(id string) (string, error)
// }

type ILoanCollections interface {
	CreateLoanTable(LoanCollections *models.LoanCollections) (string, error)
	// UpdateLoanCollections(id string) (string, error)
	// DeleteLoanCollections(LoanCollections *models.LoanCollections) (string, error)
	// GetLoanCollections(id string) (string, error)
}

// type IAccountsTable interface {
// 	CreateAccountsTable(Account *models.AccountsTable) (string, error)
// 	UpdateAccountsTable(id string) (string, error)
// 	DeleteAccountsTable(Account *models.AccountsTable) (string, error)
// 	GetAccountsTable(id string) (string, error)
// }

type IBankTable interface {
	CreateBankTable(Bank *models.BankTable) (string, error)
	// UpdateBankTable(id string) (string, error)
	// DeleteBankTable(Bank *models.BankTable) (string, error)
	// GetBankTable(id string) (string, error)
}
