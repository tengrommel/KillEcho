package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func yallo(context echo.Context) error {
	return context.String(http.StatusOK, "yallo from the web side!")
}

func main()  {
	fmt.Println("Welcome to the server")
	e := echo.New()
	e.GET("/", yallo)
	e.Start(":8000")
}