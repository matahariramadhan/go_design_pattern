package abstractfactory

import "fmt"

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
