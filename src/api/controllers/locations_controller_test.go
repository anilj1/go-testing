package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"api/domain/location"
	"api/services"
	"api/utils/error"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// We are hosting the GetCountry function in the mock. This is the function
// we are going to change on demand.
// Since this is a packege level variable is of type function,
// we can implement it as many times and as many types as.
var (
	getCountryFunc func(countryId string) (*location.Country, *error.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

// This is the mocked struct for locations service.
type locationsServiceMock struct {}

// This is implementing the locationsService interface (GetCountry) through the mock.
func (*locationsServiceMock) GetCountry(countryId string) (*location.Country, *error.ApiError) {
	// We are just returning the getCountryFunc variable declared above.
	return getCountryFunc(countryId)
}

func TestGetCountryNotFoundMockedHttp(t *testing.T) {

	// We are calling the mocked interface and not the REST APIs.
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:            "https://api.mercadolibre.com/countries/ARM",
		HTTPMethod:     http.MethodGet,
		RespHTTPCode:   http.StatusNotFound,
		RespBody:       `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params {
		{Key: "country_id", Value: "ARM"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiError error.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "Country not found", apiError.Message)
}

func TestGetCountryNotFoundMockedInterface(t *testing.T) {

	// We are calling the mocked interface and not the REST APIs.
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:            "https://api.mercadolibre.com/countries/ARM",
	// 	HTTPMethod:     http.MethodGet,
	// 	RespHTTPCode:   http.StatusNotFound,
	// 	RespBody:       `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	// })

	// We provide the definition to the getCountryFunc variables defined in this package.
	// This way, we can remove the REST level mocks as we do not need it anymore.
	getCountryFunc = func(countryId string) (country *location.Country, apiError *error.ApiError) {
		return nil, &error.ApiError {
			Status: http.StatusNotFound,
			Message:"Country not found",
		}
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params {
		{Key: "country_id", Value: "ARM"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiError error.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "Country not found", apiError.Message)
}

func TestGetCountryNoError(t *testing.T) {

	// We provide the definition to the getCountryFunc variables defined in this package.
	// This way, we can remove the REST level mocks as we do not need it anymore.
	getCountryFunc = func(countryId string) (country *location.Country, apiError *error.ApiError) {
		return &location.Country{Id:"AR", Name:"Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params {
		{Key: "country_id", Value: "ARM"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country location.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}

