package chain_of_responsibility

import (
	"fmt"
)

type Reception struct {
	next Department
}

func (r *Reception) Execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient registration already done")
	} else {
		fmt.Println("Reception registering patient")

		patient.registrationDone = true
	}

	r.next.Execute(patient)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}
