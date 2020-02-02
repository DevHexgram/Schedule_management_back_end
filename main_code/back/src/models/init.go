package models

func (s *Service) Init() {
	s.ConfigInit()
	s.DBInit()
	//s.RouterInit()
}
