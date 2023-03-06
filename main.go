package main

import (
	myrouter "project/router"
	"project/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := myrouter.Router()
	r.Run()
}
