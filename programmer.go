package main

type programmer struct {
}

func (p *programmer) makeDesk() iDesk {
	return &programmerDesk{
		desk: desk{
			deskType: "standing desk",
			size:     "dual monitor size",
		},
	}
}

func (p *programmer) makePainting() iPainting {
	return &programmerPainting{
		painting: painting{
			image: "Frodo Baggins",
		},
	}
}

func (p *programmer) makePlant() iPlant {
	return &programmerPlant{
		plant: plant{
			plantType: "cactus",
		},
	}
}
