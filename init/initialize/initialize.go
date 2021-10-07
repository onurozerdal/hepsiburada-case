package initialize

import (
	cc "hepsiburada-case/init/config"
	"hepsiburada-case/init/database"
	"hepsiburada-case/init/infrastructure"
	"hepsiburada-case/init/repository"
)

var config cc.Config

func Initialize() {
	config.Read()

	connection := database.NewConnection()
	sqlHandler := infrastructure.NewSqlHandler(connection)
	repositoryCreate := repository.NewRepository(*sqlHandler)
	repositoryCreate.CreateTable()

	connectionDb := database.NewDbConnection()
	sqlHandlerDb := infrastructure.NewSqlHandler(connectionDb)
	repositoryDb := repository.NewRepository(*sqlHandlerDb)
	repositoryDb.CreateTable()
	repositoryDb.CreateHistory()
	repositoryDb.CreateBestseller()
	repositoryDb.CreateProducts()
}