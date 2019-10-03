package main

import (
	"fmt"
)

// DataTransferObject interface to be implemented by other flyweight
// objects
type DataTransferObject interface {
	getID() string
}

// DataTransferObjectFactory struct to keep a pool
// of created objects
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}


// getDataTransferObject method that will check if the required
// object is in the pool, return it or create a new one if the
// object does not exist in the pool
func (factory *DataTransferObjectFactory) getDataTransferObject(dataObjectType string) DataTransferObject {
	dataObject, ok := factory.pool[dataObjectType]

	if !ok {
		fmt.Println("new dataObject to be created " + dataObjectType)

		switch dataObjectType {
		case "customer":
			dataObject = &Customer{id: "1"}
			factory.pool[dataObjectType] = dataObject
		case "employee":
			dataObject = &Employee{id: "2"}
			factory.pool[dataObjectType] = dataObject
		case "manager":
			dataObject = &Manager{id: "3"}
			factory.pool[dataObjectType] = dataObject
		case "address":
			dataObject = &Address{id: "4"}
			factory.pool[dataObjectType] = dataObject
		default:
			// id can be randomly generated
			dataObject = &CustomObject{id: "5", name: dataObjectType}
			factory.pool[dataObjectType] = dataObject
		}
	}

	return dataObject
}

// Customer model
type Customer struct {
	id   string // sequence generator
	name string
	ssn  string
}

// Customer getID method that satisfies the DataTransferObject
// flyweight interface
func (customer *Customer) getID() string {
	return customer.id
}

// Employee flyweight
type Employee struct {
	id   string
	name string
}

// Employee getID method
func (employee *Employee) getID() string {
	return employee.id
}

// Manager flyweight struct
type Manager struct {
	id         string
	name       string
	department string
}

// Manager getID method
func (manager *Manager) getID() string {
	return manager.id
}

// Address flyweight struct
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	city        string
	state       string
	country     string
}

// Address getID method
func (address *Address) getID() string {
	return address.id
}

// CustomObject used to create any object that is not in the pool
type CustomObject struct {
	id string
	name string
}

// CustomObject getID method
func (custom *CustomObject) getID() string {
	return custom.id
}

func main() {
	factory := DataTransferObjectFactory{
		pool: make(map[string]DataTransferObject),
	}

	customer := factory.getDataTransferObject("customer")
	employee := factory.getDataTransferObject("employee")
	manager := factory.getDataTransferObject("manager")
	address := factory.getDataTransferObject("address")
	// create new flyweight objects
	occupation := factory.getDataTransferObject("occupation")
	education := factory.getDataTransferObject("education")

	fmt.Println("Customer id " + customer.getID())
	fmt.Println("Employee id " + employee.getID())
	fmt.Println("Manager id " + manager.getID())
	fmt.Println("Address id " + address.getID())
	fmt.Println("Occupation id " + occupation.getID())
	fmt.Println("Education id " + education.getID())
}