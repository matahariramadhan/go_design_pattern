package builder

import "fmt"

func Main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	normalDirector := newDirector(normalBuilder)
	normalHouse := normalDirector.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	iglooDirector := newDirector(iglooBuilder)
	iglooHouse := iglooDirector.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
	// printHouseDetails(igloHouse)
}
