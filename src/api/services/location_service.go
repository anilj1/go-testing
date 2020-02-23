package services

import (
	"fmt"

	"api/domain/location"
	"api/provider/locprovider"
	"api/utils/error"
)

type locationsService struct {}

//
type locationsServiceInterface interface {
	GetCountry(countryId string) (*location.Country, *error.ApiError)
}

// Instead of making the public variable (LocationsService) an instance of
// pointer of locationsService struct, we make it a variable of interface.
// Any struct implementing this interface can be a locationsService.
var (
	LocationsService locationsServiceInterface
)


// Some details about init function:
// Every package in go has init() function. When a package is imported,
// it asks if the package has been initialized before. If the flag is 'false'
// it calls the init() function on that package

// In order for locationsService to implement locationsServiceInterface, we
// need to have all the methods in interface implemented by locationService.
func init() {
	LocationsService = &locationsService{}
}

// There is a reason to use service than just the location provider.
// We are not able to mock a function inside a package. Since go is a
// compiled language, once the test case is compiled, the function is also
// compiled and test case is executed againset he compiled code.

// The only way is to mock the function on a struct. The function GetCountry()
// becomes the function of the struct.
func (s *locationsService) GetCountry(countryId string) (*location.Country, *error.ApiError) {
	fmt.Println("Inside service")
	return locprovider.GetCountry(countryId)
}
