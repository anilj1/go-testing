package test

import (
	"fmt"
	"os"
	"testing"

	"api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("About to start the application...")
	go app.StartApp()
	fmt.Println("Application started, about to start functional tests...")
	os.Exit(m.Run())
}


