package main

import "fmt"

func main() {
	programmerFactory, _ := getOfficeDecorationFactory("programmer")
	salespersonFactory, _ := getOfficeDecorationFactory("salesperson")

	salespersonDesk := salespersonFactory.makeDesk()
	salespersonPlant := salespersonFactory.makePlant()
	salespersonPainting := salespersonFactory.makePainting()

	programmerDesk := programmerFactory.makeDesk()
	programmerPlant := programmerFactory.makePlant()
	programmerPainting := programmerFactory.makePainting()

	fmt.Println("Salesperson Desk Details:")
	printDeskDetails(salespersonDesk)
	printDeskPlant(salespersonPlant)
	printPaintingToHang(salespersonPainting)

	fmt.Println()

	fmt.Println("Programmer Desk Details:")
	printDeskDetails(programmerDesk)
	printDeskPlant(programmerPlant)
	printPaintingToHang(programmerPainting)
}

func printDeskDetails(s iDesk) {
	fmt.Printf("Desk Type: %s", s.getDeskType())
	fmt.Println()
	fmt.Printf("Desk Size: %s", s.getSize())
	fmt.Println()
}

func printDeskPlant(s iPlant) {
	fmt.Printf("Plant Type: %s", s.getPlantType())
	fmt.Println()
}

func printPaintingToHang(s iPainting) {
	fmt.Printf("Painting Image: %s", s.getImage())
	fmt.Println()
}
