package abstract_factory

import (
	"errors"
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Sport motorbike
	motorbikeVehicle, err := motorbikeFactory.Build(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Sport motorbike vehicle has %d wheels\n ", motorbikeVehicle.NumWheels())
	t.Logf("Sport motorbike vehicle has %d seats\n ", motorbikeVehicle.NumSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())

	// Cruiser motorbike
	motorbikeVehicle, err = motorbikeFactory.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Cruise motorbike vehicle has %d wheels\n ", motorbikeVehicle.NumWheels())
	t.Logf("Cruise motorbike vehicle has %d seats\n ", motorbikeVehicle.NumSeats())

	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetMotorbikeType())

	_, err = motorbikeFactory.Build(3)
	if err == nil {
		t.Fatal(errors.New("expected error, got nil"))
	}

	t.Logf("Expected error: %s", err.Error())
}
