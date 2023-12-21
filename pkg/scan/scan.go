package scan

type Scan struct {
}

func New() *Scan {
	s := &Scan{}
	return s
}

func (s *Scan) Run() error {
	return nil
}
