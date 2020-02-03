package main

import (
	_ "github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/setting"
	"github.com/DevHexgram/Schedule_management_back_end/routers"
)

func main() {
	r := routers.RouterInit()
	r.Run(setting.Config.Web.Port)
}
