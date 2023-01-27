package factorymethod

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "ak47",
			power: 4,
		},
	}
}
