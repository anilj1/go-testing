package app

import (
	"api/controllers"
)

func mapUrls() {

	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
