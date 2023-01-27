package builder

type NormalBuilder struct {
	doorType   string
	windowType string
	floor      int
}

func newNormalBuilder() IBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}
func (b *NormalBuilder) setNumFloor() {
	b.floor = 1
}
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
