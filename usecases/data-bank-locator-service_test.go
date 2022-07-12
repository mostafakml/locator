package usecases

import (
	config2 "github.com/mostafakml/locator/config"
	"github.com/mostafakml/locator/domain"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDataBankService_Calculate(t *testing.T) {
	dataBankService:=NewDataBankService()
	config:=config2.APPConfig{
		Port: "8080",
		SectorId: 1.0,
	}

	drone :=domain.Drone{
		XValue: "1",
		YValue: "1",
		ZValue: "1",
		Vel: "1",
	}

	dataBankLocation,err:=dataBankService.Calculate(drone,config)
	assert.Nil(t,err)
	assert.Equal(t,"1.000000",dataBankLocation.Location)
}

func TestDataBankService_convertToFloat(t *testing.T)  {

	testCase:="1.0"

	dataBankService:=NewDataBankService()

	f:=dataBankService.convertToFloat(testCase)
	assert.Equal(t,1.0,f)

}


func TestDataBankService_convertToString(t *testing.T)  {

	testCase:=1.0

	dataBankService:=NewDataBankService()

	f:=dataBankService.convertToString(testCase)
	assert.Equal(t,"1.000000",f)

}


