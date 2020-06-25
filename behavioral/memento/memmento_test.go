package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())

	caretaker.addMemento(originator.createMemento())
	originator.setState("C")

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restorememento(caretaker.getMenento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restorememento(caretaker.getMenento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
