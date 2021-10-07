package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hepsiburada-case/api/database"
	"hepsiburada-case/api/endpoint"
	"hepsiburada-case/api/infrastructure"
	"hepsiburada-case/api/repository"
	"hepsiburada-case/api/service"
)

var apiController *endpoint.ApiController

func init() {
	connection := database.NewRecommendedDbConnection()
	sqlHandler := infrastructure.NewSqlHandler(connection)
	apiRepository := repository.NewApiRepository(*sqlHandler)
	apiService := service.NewApiService(*apiRepository)
	apiController = endpoint.NewApiController(*apiService)
}

func main() {
	e := echo.New()
	e.GET("/api/browsingHistories", apiController.BrowsingHistories)
	e.GET("/api/bestsellerProducts", apiController.BestsellerProducts)
	e.DELETE("/api/deleteHistory", apiController.DeleteHistory)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "8080")))
}
