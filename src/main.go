package main

import (
	"fmt"
	"github.com/DevHexgram/Schedule_management_back_end/models"
	"github.com/DevHexgram/Schedule_management_back_end/pkg/setting"
)

func main() {
	fmt.Println(models.DB)
	fmt.Println(setting.Config)
}
