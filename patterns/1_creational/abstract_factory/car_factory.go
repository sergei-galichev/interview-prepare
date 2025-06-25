package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = iota + 1
	FamilyCarType
)

type CarFactory struct{}

func (f *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
