package main

import (
	"fmt"
)

// IProcess interface
type IProcess interface {
	process() string
}

// Process type for creating processes
type Process struct {
	name string
}

// process method to implemented for the IProcess interface
func (process *Process) process() string {
	return ("Process " + process.name)
}

// ProcessDecorator (wrapper) to add more functionality to process
// uses the field processInstance to link to the wrapper
type ProcessDecorator struct {
	processInstance *Process
}

// process method satisfied for the IProcess interface
func(decorator *ProcessDecorator) process() string {
	return ("Process Decorator with wrappee method " + decorator.processInstance.process())
}

// additional method to add more functionality to process instance
func(decorator *ProcessDecorator) changeState() string {
	return "state of process changed from down to up"
}

func main() {
	var process = &Process{
		name: "string",
	}

	var decorator = &ProcessDecorator{
		processInstance: process,
	}

	fmt.Println(decorator.changeState())
	fmt.Println(decorator.process())
}
