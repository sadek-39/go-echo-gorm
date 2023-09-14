package config

import "fmt"

const (
	DBUser     = "user"
	DBPassword = "user"
	DBName     = "go-echo-db"
	DBHost     = "localhost"
	DBPort     = "3307"
	DBType     = "mysql"
)

func GetDBType() string {
	return DBType
}

func GetMYSQLConnectionString() string {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)
	return database
}
