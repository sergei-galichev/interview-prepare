package chain_of_responsibility

import (
	"fmt"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("Payment done")
	}

	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
