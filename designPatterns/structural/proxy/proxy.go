package main

import (
	"fmt"
)

// IRealObject Interface that the real object and proxy object
// has to implement
type IRealObject interface {
	performAction()
}

// RealObject struct that would be the original service
type RealObject struct {}

// the real object implements the IRealObject's Interface
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

// VirtualProxy struct that will stand in place for the realObject
// it will have a field that instantiate with the realObject service
// where the real operation for the realObject service are forwarded
type VirtualProxy struct {
	realService *RealObject
}

// VirtualProxy performAction method to implement the IRealObject
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realService == nil {
		virtualProxy.realService = &RealObject{}
	}

	fmt.Println("Proxy action performAction()")
	virtualProxy.realService.performAction()
}

func main() {
	// the object client has the interface
	var object IRealObject = &VirtualProxy{}
	object.performAction()
}