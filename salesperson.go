package main

type salesperson struct {
}

func (s *salesperson) makeDesk() iDesk {
	return &salespersonDesk{
		desk: desk{
			size:     "laptop surface",
			deskType: "collapsible table",
		},
	}
}

func (s *salesperson) makePainting() iPainting {
	return &salespersonPainting{
		painting: painting{
			image: "sleepy kittens",
		},
	}
}

func (s *salesperson) makePlant() iPlant {
	return &salespersonPlant{
		plant: plant{
			plantType: "lilly",
		},
	}
}
