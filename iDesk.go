package main

type iDesk interface {
	setDeskType(deskType string)
	setSize(size string)
	getDeskType() string
	getSize() string
}

type desk struct {
	deskType string
	size     string
}

func (s *desk) setDeskType(deskType string) {
	s.deskType = deskType
}

func (s *desk) getDeskType() string {
	return s.deskType
}

func (s *desk) setSize(size string) {
	s.size = size
}

func (s *desk) getSize() string {
	return s.size
}
