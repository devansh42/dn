package main

import (
	"dn/page"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("hello", page.MainPage)

	e.Start(":8080")
}
func handler(e echo.Context) error {
	return e.HTML(200, "Hey , My name is Devansh Gupta")
}

//https://ssh.cloud.google.com/projects/stark-218116/zones/asia-south1-c/instances/instance-1?authuser=0&hl=en_US&projectNumber=137013794350
