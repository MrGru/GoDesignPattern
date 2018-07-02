package abstract_factory

type CruiseMotobike struct{}

func (c *CruiseMotobike) NumWheels() int {
	return 2
}

func (c *CruiseMotobike) NumSeats() int {
	return 2
}

func (c *CruiseMotobike) GetMotorbikeType() int {
	return CruiseMotobikeType
}
