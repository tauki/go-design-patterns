package singleton

type Singleton interface {
	AddOne() int
	GetCount() int
}

type singleton struct {
	count int
}

var instance *singleton

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

func (s *singleton) GetCount() int {
	return s.count
}
