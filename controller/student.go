package controller

import (
	"github.com/labstack/echo/v4"
	"go-with-echo-gorm/model"
	"go-with-echo-gorm/storage"
	"net/http"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetClass(c echo.Context) error {
	class, _ := GETRepoClass()

	return c.JSON(http.StatusOK, class)

}

func GetRepoStudents() ([]model.Student, error) {
	db := storage.GetDBInstance()

	studentlist := []model.Student{}

	if err := db.Find(&studentlist).Error; err != nil {
		return nil, err
	}
	return studentlist, nil
}

func GETRepoClass() ([]model.Class, error) {
	db := storage.GetDBInstance()

	class := []model.Class{}

	err := db.Find(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}
