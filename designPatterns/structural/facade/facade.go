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
