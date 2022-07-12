package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mostafakml/locator/config"
	"github.com/mostafakml/locator/domain"
	"github.com/mostafakml/locator/interfaces/request"
	"github.com/mostafakml/locator/usecases"
	"net/http"
)


type DataBankLocationController struct {
	DataBankService usecases.DataBankService
	Config config.APPConfig
}
var validate *validator.Validate


func NewDataBankLocationController(DataBankService usecases.DataBankService,Config config.APPConfig) DataBankLocationController {
	validate=validator.New()
	validate.RegisterValidation("is_convertible",request.ValidateIsConvertible)

	return DataBankLocationController{
		DataBankService: DataBankService,
		Config: Config,
	}

}

func (h DataBankLocationController) Locate(c *gin.Context)  {
	b := domain.Drone{}


	if err := c.ShouldBindJSON(&b); err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	err:=validate.Struct(&b)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	g, err := h.DataBankService.Calculate(b,h.Config)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{"data": g})


}


