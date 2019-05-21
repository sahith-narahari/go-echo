package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from echo1")
}

func getCats(c echo.Context) error {
	dataObtained := c.Param("data")
	catname := c.QueryParam("name")
	cattype := c.QueryParam("type")

	if dataObtained == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s \n and his type is: %s \n", catname, cattype))
	}
	if dataObtained == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catname,
			"type": cattype,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "you need to give proper data - json or string",
		})
	}
}

func main() {
	fmt.Println("Hello World!")

	e := echo.New()

	e.GET("/", hello)

	e.GET("/cats/:data", getCats)
	e.Logger.Fatal(e.Start(":3000"))
}
