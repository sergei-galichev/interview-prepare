package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType = iota + 1
	MotorbikeFactoryType
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
