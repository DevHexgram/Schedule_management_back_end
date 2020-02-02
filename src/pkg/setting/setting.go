package setting

import (
	"github.com/BurntSushi/toml"
)

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

var Config Conf

func init() {
	//confPath := "../SMS_config/conf.toml"

	confPath := "conf/conf.toml"


	_, err := toml.DecodeFile(confPath, &Config)

	if err != nil {
		panic(err)
	}
	//fmt.Println(s.Config)

	//s.Config = *c
}
