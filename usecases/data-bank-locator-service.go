package usecases

import (
	"fmt"
	"github.com/mostafakml/locator/config"
	"github.com/mostafakml/locator/domain"
	"strconv"
)

type DataBankService struct {
}

var sss float64


func NewDataBankService() DataBankService {
	return DataBankService{}
}


func (dataBankService DataBankService) Calculate(drone domain.Drone,config config.APPConfig) (domain.DataBankLocation,error) {


	println()

	Floatlocation:=dataBankService.convertToFloat(drone.XValue)*dataBankService.convertToFloat(drone.YValue)*dataBankService.convertToFloat(drone.ZValue)*dataBankService.convertToFloat(drone.Vel)*config.SectorId

	strLocation:=dataBankService.convertToString(Floatlocation)
	location:=domain.DataBankLocation{
		Location: strLocation,
	}
	return location ,nil
}


func (dataBankService DataBankService) convertToFloat(str string ) float64 {

	floatVal ,err :=strconv.ParseFloat(str,64)

	if err!=nil {
		panic(err)
	}

	return floatVal
}

func (dataBankService DataBankService) convertToString(floatVal float64 ) string {

	s := fmt.Sprintf("%f", floatVal)

	return s
}
