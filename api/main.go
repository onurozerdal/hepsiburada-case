package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/onurozerdal/hepsiburada-case/api/database"
	"github.com/onurozerdal/hepsiburada-case/api/endpoint"
	"github.com/onurozerdal/hepsiburada-case/api/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/api/repository"
	"github.com/onurozerdal/hepsiburada-case/api/service"
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
