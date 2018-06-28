package creational

//Singleton interface for count number increasement
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

//GetInstance for get singleton object
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
