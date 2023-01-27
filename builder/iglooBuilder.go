package builder

type IglooBuilder struct {
	doorType   string
	windowType string
	floor      int
}

func newIglooBuilder() IBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		doorType:   i.doorType,
		windowType: i.windowType,
		floor:      i.floor,
	}
}
