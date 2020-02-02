package models

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Config Conf
	DB     *gorm.DB
	Router *gin.Engine
}

type DBConfig struct {
	Address  string
	User     string
	Password string
	DBName   string
}

type webConfig struct {
	Port string
}

type Conf struct {
	DB  DBConfig
	Web webConfig
}

func (s *Service) ConfigInit() {
	//confPath := "../SMS_config/conf.toml"

	confPath := "./SMS_config/conf_test.toml"
	//c := new(Conf)

	_, err := toml.DecodeFile(confPath, &s.Config)

	if err != nil {
		panic(err)
	}
	//fmt.Println(s.Config)

	//s.Config = *c
}
