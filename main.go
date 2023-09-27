package main

import (
	"github.com/labstack/echo/v4"
	"go-with-echo-gorm/controller"
	"go-with-echo-gorm/storage"
	"net/http"
)

func main() {
	e := echo.New()
	storage.NewDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/students", controller.GetStudents)
	e.GET("/class", controller.GetClass)

	e.Logger.Fatal(e.Start(":1323"))
}
