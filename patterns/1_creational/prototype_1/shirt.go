package prototype_1

import (
	"fmt"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt: %s, price: %f, color: %d", s.SKU, s.Price, s.Color)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var (
	whitePrototype = &Shirt{
		Price: 15.0,
		SKU:   "empty",
		Color: White,
	}

	blackPrototype = &Shirt{
		Price: 16.0,
		SKU:   "empty",
		Color: Black,
	}

	bluePrototype = &Shirt{
		Price: 17.0,
		SKU:   "empty",
		Color: Blue,
	}
)
