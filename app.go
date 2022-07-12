package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mostafakml/locator/config"
	"github.com/mostafakml/locator/interfaces/controller"
	"github.com/mostafakml/locator/usecases"
)
var appConfig = config.APPConfig{}

func main() {
	appConfig.Read()
	setupRouter()

}
func setupRouter()  {
	r := gin.Default()
	r1 := controller.NewDataBankLocationController (usecases.NewDataBankService(), appConfig)

	r.POST("/v1/locate", r1.Locate)

	r.Run(":"+appConfig.Port)
}