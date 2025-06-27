package prototype_1

import (
	"errors"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = iota + 1
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct{}

func (sc *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
		//return whitePrototype, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
