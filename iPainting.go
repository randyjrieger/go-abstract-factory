package main

type iPainting interface {
	setImage(image string)
	getImage() string
}

type painting struct {
	image string
}

func (s *painting) setImage(image string) {
	s.image = image
}

func (s *painting) getImage() string {
	return s.image
}
