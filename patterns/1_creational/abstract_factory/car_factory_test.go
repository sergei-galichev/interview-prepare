package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Luxury car
	carVehicle, err := carFactory.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Luxury car vehicle has %d wheels\n ", carVehicle.NumWheels())
	t.Logf("Luxury car vehicle has %d seats\n ", carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors\n", luxuryCar.NumDoors())

	// Family car
	familyVehicle, err := carFactory.Build(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Family car vehicle has %d wheels\n ", familyVehicle.NumWheels())
	t.Logf("Family car vehicle has %d seats\n ", familyVehicle.NumSeats())

	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Family car has %d doors\n", familyCar.NumDoors())

	_, err = carFactory.Build(3)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
