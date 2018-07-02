package abstract_factory

type SportMotobike struct{}

func (s *SportMotobike) NumWheels() int {
	return 2
}

func (s *SportMotobike) NumSeats() int {
	return 1
}

func (s *SportMotobike) GetMotorbikeType() int {
	return SportMotobikeType
}
