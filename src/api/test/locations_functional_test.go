package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	error2 "api/utils/error"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {

	// Mockup
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:            "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:     http.MethodGet,
		RespHTTPCode:   http.StatusNotFound,
		RespBody:       `{"status":404, "error":"not_found", "message":"no country with id AR"}`,
	})

	// Execution
	fmt.Println("About to test get countries not found test")
	response, err := http.Get("http://localhost:8080/locations/countries/AR")

	// Validation 
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Body: ", string(bytes))

	var apiErr error2.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id AR", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T) {

	// Mockup
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:            "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:     http.MethodGet,
		RespHTTPCode:   http.StatusOK,
		RespBody:       `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}},"states":[{"id":"AR-B","name":"Buenos Aires"},{"id":"AR-C","name":"Capital Federal"},{"id":"AR-K","name":"Catamarca"},{"id":"AR-H","name":"Chaco"},{"id":"AR-U","name":"Chubut"},{"id":"AR-W","name":"Corrientes"},{"id":"AR-X","name":"Córdoba"},{"id":"AR-E","name":"Entre Ríos"},{"id":"AR-P","name":"Formosa"},{"id":"AR-Y","name":"Jujuy"},{"id":"AR-L","name":"La Pampa"},{"id":"AR-F","name":"La Rioja"},{"id":"AR-M","name":"Mendoza"},{"id":"AR-N","name":"Misiones"},{"id":"AR-Q","name":"Neuquén"},{"id":"AR-R","name":"Río Negro"},{"id":"AR-A","name":"Salta"},{"id":"AR-J","name":"San Juan"},{"id":"AR-D","name":"San Luis"},{"id":"AR-Z","name":"Santa Cruz"},{"id":"AR-S","name":"Santa Fe"},{"id":"AR-G","name":"Santiago del Estero"},{"id":"AR-V","name":"Tierra del Fuego"},{"id":"AR-T","name":"Tucumán"}]}`,
	})

	fmt.Println("About to test get countries not found test")
	response, err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Body: ", string(bytes))
}
