package main

import (
	"fmt"

	"api/app"
	"api/provider/locprovider"
)

func main() {
	country, err := locprovider.GetCountry("AR")
	fmt.Println(err)
	fmt.Println(country)

	app.StartApp()
}
