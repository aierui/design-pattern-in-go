package responsibility

import "fmt"

type cashier struct {
	next department
}

func (m *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (m *cashier) setNext(next department) {
	m.next = next
}
