package main

import (
	"fmt"
)

// Account model
type Account struct {
	id          string
	accountType string
}

// Account method to create an account type
func (account *Account) create(accountType string) *Account {
	fmt.Println("Account creation with type")

	account.accountType = accountType
	return account
}

// Method to get an account by its id
func(account *Account) getByID(id string) *Account {
	fmt.Println("Get an account by Id " + id)

	// make a query to the db and return the data
	return account
}

// Method needed to delete an account with a given valid id
func(account *Account) deleteByID(id string) {
	fmt.Printf("Delete account with ID %s", id)
}

// Customer model
type Customer struct {
	id int
	name string
}

// customer has a method to create customer with a name
func(customer *Customer) create(name string) *Customer {
	fmt.Println("Creating customer...")

	customer.name = name
	return customer
}

// Transaction model
type Transaction struct {
	id string
	amount float32
	srcAccountID string
	destAccountID string
}

// Transaction has a create method to help create a transaction
func(transaction *Transaction) create(srcAccountID, destAccountID string, amount float32) *Transaction {
	fmt.Println("creating a transaction")

	transaction.srcAccountID = srcAccountID
	transaction.destAccountID = destAccountID
	transaction.amount = amount

	return transaction
}

// BranchManagerFacade to access the different services
type BranchManagerFacade struct {
	account *Account
	customer *Customer
	transaction *Transaction
}

// newBranchManager function to create a new BranchManagerFacade
func newBranchManager() *BranchManagerFacade {
	branchManagerFacade := &BranchManagerFacade{
		&Account{},
		&Customer{},
		&Transaction{},
	}

	return branchManagerFacade
}

// method for branch manager to create a customer account
func(branchManager *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
	customer := branchManager.customer.create(customerName)
	account := branchManager.account.create(accountType)

	return customer, account
}

// method for the new branch manager to perform a transaction
func(branchManager *BranchManagerFacade) createTransaction(srcAccountID, destAccountID string, amount float32) *Transaction {
	transaction := branchManager.transaction.create(srcAccountID, destAccountID, amount)

	return transaction
}

func main()  {
	// create the new branch manager facade
	branchManager := newBranchManager()

	// create a customer account
	customer, account := branchManager.createCustomerAccount("Olumide Ogundele", "Current Account")

	fmt.Printf("Customer name: %s \n", customer.name)
	fmt.Printf("Customer's account type: %s \n", account.accountType)

	// perform a transaction
	transaction := branchManager.createTransaction("45678", "98576", 200000.56)

	fmt.Printf("Current transaction amount: %g \n", transaction.amount)
}