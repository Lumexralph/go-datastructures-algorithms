package main

import (
	"fmt"
	"encoding/json"
)

type AccountDetails struct {
	id          string
	accountType string
}

// Account struct  having 2 fields CustomerName is made
// public by exporting it while details is made private
// by not making it accessible in the  struct
type Account struct {
	CustomerName string
	details      *AccountDetails
}

// method to set an account details
func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

// method to get an account by id
func (account *Account) getID() string {
	return account.details.id
}

// method to return account type
func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	account := &Account{CustomerName: "Olumide Ogundele"}
	account.setDetails("6309", "savings")

	// converting the account to a JSON will ignore the private fields
	// i.e fields that were not exported (starting with lowercase)
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Display Account without hidden account details", string(jsonAccount))
	fmt.Println("Account ID ", account.getID())
	fmt.Println("Account Type ", account.getAccountType())
}
