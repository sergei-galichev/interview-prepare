package chain_of_responsibility

import (
	"fmt"
)

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Doctor checking patient")

		patient.doctorCheckUpDone = true
	}

	d.next.Execute(patient)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
