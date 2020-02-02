//初始化serve对象
package main

import "main/models"

func (s *models.Service) Init() {
	s.ConfigInit()
	s.DBInit()
	s.RouterInit()
}

func main()  {
	var serve models.Service
	serve.Init()
}

