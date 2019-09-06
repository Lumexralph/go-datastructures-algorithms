package main

import "fmt"

/*Let's say you have an IProcessor interface with a
process method, the Adapter class implements the
process method and has an Adaptee instance as an attribute.
The Adaptee class has a convert method and
an adapterType instance variable. The developer while using
the API client calls the process interface method to invoke
convert on Adaptee.
*/

// IProcess interface
type IProcess interface {
	process()
}

// Adapter struct - form of class in Python or Java
type Adapter struct {
	adaptee Adaptee
}

// Adapter struct/class process method
func (adapter Adapter) process() {
	fmt.Println("Adapter process")

	// call the convert method on the adaptee
	adapter.adaptee.convert()
}

// Adaptee struct/class
type Adaptee struct {
	adapterType int
}

// Adaptee struct convert method
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method in action!")
}


// start the whole program
func main() {
	// the Adapter implements the IProcessor's interface
	// IProcess can be a legacy code, third party API
	// the adapter instance datatype can be assigned to the
	// interface, the client will interface with this interface
	var processor IProcess = Adapter{}
	processor.process()
}