package abstractfactory

type Nike struct {
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 12,
		},
	}
}
