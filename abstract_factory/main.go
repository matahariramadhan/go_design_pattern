package abstractfactory

import "fmt"

func Main() {
	nikeFactory, _ := GetSportsFactory("nike")
	adidasFactory, _ := GetSportsFactory("adidas")

	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
