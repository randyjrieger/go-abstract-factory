package main

import "fmt"

type iOfficeDecorateFactory interface {
	makeDesk() iDesk
	makePainting() iPainting
	makePlant() iPlant
}

func getOfficeDecorationFactory(positionType string) (iOfficeDecorateFactory, error) {
	if positionType == "programmer" {
		return &programmer{}, nil
	}

	if positionType == "salesperson" {
		return &salesperson{}, nil
	}

	return nil, fmt.Errorf("Invalid position type")
}
