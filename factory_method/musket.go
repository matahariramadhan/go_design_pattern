package factorymethod

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}
