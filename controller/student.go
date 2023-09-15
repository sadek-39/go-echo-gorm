package controller

import (
	"github.com/labstack/echo/v4"
	"go-with-echo-gorm/model"
	"go-with-echo-gorm/storage"
	"log"
	"net/http"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Student, error) {
	db := storage.GetDBInstance()
	log.Print(db)
	students := []model.Student{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
