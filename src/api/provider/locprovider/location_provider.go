package locprovider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api/domain/location"
	"api/utils/error"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	getCountryUrl = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*location.Country, *error.ApiError) {

	response := rest.Get(fmt.Sprintf(getCountryUrl, countryId))
	if response == nil || response.Response == nil {
		return nil, &error.ApiError {
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		var apiError error.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiError); err != nil {
			return nil, &error.ApiError {
				Status: http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiError
	}

	var result location.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &error.ApiError {
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &result, nil
}
