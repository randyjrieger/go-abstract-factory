package main

type iPlant interface {
	setPlantType(plantType string)
	getPlantType() string
}

type plant struct {
	plantType string
}

func (s *plant) setPlantType(plantType string) {
	s.plantType = plantType
}

func (s *plant) getPlantType() string {
	return s.plantType
}
