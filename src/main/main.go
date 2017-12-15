package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func yallo(context echo.Context) error {
	return context.String(http.StatusOK, "yallo from the web side!")
}

func getCats(context echo.Context) error {
	catName := context.QueryParam("name")
	catType := context.QueryParam("type")

	dataType := context.Param("data")

	if dataType == "string"{
		return context.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand his type is: %s\n", catName, catType))
	}
	if dataType == "json"{
		return context.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return context.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})
}

func main()  {
	fmt.Println("Welcome to the server")
	e := echo.New()

	e.GET("/", yallo)
	e.GET("/cats/:data", getCats)

	e.Start(":8000")
}