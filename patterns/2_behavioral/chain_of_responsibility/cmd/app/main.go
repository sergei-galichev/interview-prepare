package main

import (
	"chain_of_responsibility"
)

func main() {
	cashier := &chain_of_responsibility.Cashier{}

	medical := &chain_of_responsibility.Medical{}
	medical.SetNext(cashier)

	doctor := &chain_of_responsibility.Doctor{}
	doctor.SetNext(medical)

	reception := &chain_of_responsibility.Reception{}
	reception.SetNext(doctor)

	patient := &chain_of_responsibility.Patient{Name: "John Doe"}

	reception.Execute(patient)
}
