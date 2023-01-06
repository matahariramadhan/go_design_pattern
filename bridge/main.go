package bridge

import "fmt"

func Main() {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	winComputer := &Windows{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

}
