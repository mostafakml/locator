package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mostafakml/locator/config"
	"github.com/mostafakml/locator/usecases"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataBankLocationController_Locate(t *testing.T) {


	service:=usecases.NewDataBankService();
	appConfig:=config.APPConfig{
		Port: "8080",
		SectorId: 1.0,
	}


	dataBankController:=NewDataBankLocationController(service,appConfig)
	r := SetUpRouter()
	r.POST("/v1/locate", dataBankController.Locate)



	testCases := []struct {

		name      string
		body      []byte
		responseBody      string
		watedCode int
	}{
		{
			name:      "when the request is valid",
			body:      []byte(`{"x_value":"1","y_value":"1","z_value":"1","vel":"1"}`),
			responseBody: `{"data":{"location":"1.000000"}}`,
			watedCode: 200,
		},
		{
			name:      "when the request is not valid",
			body:      []byte(`{"x_value":"p","y_value":"1","z_value":"1","vel":"1"}`),
			responseBody: `{"error":"Key: 'Drone.XValue' Error:Field validation for 'XValue' failed on the 'is_convertible' tag"}`,
			watedCode: 422,
		},
		{
			name:      "when the request is not valid",
			body:      []byte(`{"x_value":"1","y_value":"1.p","z_value":"1","vel":"1"}`),
			responseBody: `{"error":"Key: 'Drone.YValue' Error:Field validation for 'YValue' failed on the 'is_convertible' tag"}`,
			watedCode: 422,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {


			//rq := httptest.NewRequest(http.MethodPost, "/v1/games", strings.NewReader(tc.body))
			req, _ := http.NewRequest("POST", "/v1/locate", bytes.NewBuffer(tc.body))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			responseData, _ := ioutil.ReadAll(w.Body)
			assert.Equal(t, tc.responseBody, string(responseData))
			assert.Equal(t, tc.watedCode, w.Code)

		})
	}






}

func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}